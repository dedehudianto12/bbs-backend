package database

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/dedehudianto12/bbs-backend/config"
	"github.com/golang-migrate/migrate/v4"
	// Registers the "postgres" database driver and the "file" source driver with
	// migrate's registry. Blank imports: they are wired up by init(), not called.
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate applies any pending migrations, and is called during boot before the
// server accepts its first request.
//
// ── Why this runs on boot ───────────────────────────────────────────────────
//
// The schema used to advance only when someone ran `migrate up` from a laptop.
// Nothing tied that to a deploy, so a push could put a new binary in front of
// an old schema — and the binary would start serving immediately, querying
// columns that did not exist yet. Migration 000009 sat unapplied in production
// through a full deploy for exactly this reason; it was harmless only because
// no Go code depended on it. Running migrations here makes code and schema ship
// as one unit, because they are both in the image and both advance at startup.
//
// ── Concurrency ─────────────────────────────────────────────────────────────
//
// Safe with multiple instances. golang-migrate takes a Postgres advisory lock
// before touching anything, so if two containers boot together one applies the
// migrations and the other blocks until it can confirm there is nothing left to
// do. See MigrationDSN for why that lock forces migrations onto the direct
// endpoint rather than the pooled one.
//
// ── Failure policy ──────────────────────────────────────────────────────────
//
// Every error here is fatal to startup, deliberately. A process that cannot
// bring the schema to the version its code expects has no safe way to serve
// traffic — it would answer requests with errors from half-migrated tables. A
// container that refuses to start is louder, is caught by Railway's health
// check, and leaves the previous working deployment in place.
func Migrate(cfg *config.Config) error {
	// migrate wants a file:// URL, and a relative path in one ("file://migrations")
	// parses the first segment as a *hostname*, silently looking in the wrong
	// place. Resolving to an absolute path first avoids that entirely.
	absPath, err := filepath.Abs(cfg.Database.MigrationsPath)
	if err != nil {
		return fmt.Errorf("resolve migrations path %q: %w", cfg.Database.MigrationsPath, err)
	}
	if _, err := os.Stat(absPath); err != nil {
		return fmt.Errorf(
			"migrations directory %q is not readable: %w (set MIGRATIONS_PATH if it lives elsewhere)",
			absPath, err,
		)
	}

	m, err := migrate.New("file://"+absPath, MigrationDSN(cfg))
	if err != nil {
		return fmt.Errorf("open migrator: %w", err)
	}
	defer func() {
		// Close reports the source and database errors separately; neither is
		// worth failing an already-successful boot over, but a leaked connection
		// that never gets logged is worth avoiding.
		if srcErr, dbErr := m.Close(); srcErr != nil || dbErr != nil {
			slog.Warn("migrator close", "source", srcErr, "database", dbErr)
		}
	}()

	before, dirty, err := m.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		return fmt.Errorf("read schema version: %w", err)
	}

	// A dirty schema means a previous migration died partway through, so the
	// database is in a state no migration file describes. Guessing forward from
	// there risks compounding the damage — this needs a human who can look at
	// what actually landed, repair it, and `migrate force <version>`.
	if dirty {
		return fmt.Errorf(
			"schema is dirty at version %d: an earlier migration failed partway "+
				"through and the database is in an unknown state. Inspect it, repair "+
				"it by hand, then run `migrate force %d`. Refusing to start",
			before, before,
		)
	}

	switch err := m.Up(); {
	case errors.Is(err, migrate.ErrNoChange):
		slog.Info("schema already up to date", "version", before)
		return nil
	case err != nil:
		return fmt.Errorf("apply migrations: %w", err)
	}

	after, _, err := m.Version()
	if err != nil {
		return fmt.Errorf("read schema version after migrating: %w", err)
	}
	slog.Info("schema migrated", "from", before, "to", after)
	return nil
}
