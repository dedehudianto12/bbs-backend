package database

import (
	"net"
	"net/url"
	"strings"

	"github.com/dedehudianto12/bbs-backend/config"
)

// DSN renders the connection string the application pool uses.
//
// Built through net/url rather than fmt.Sprintf so the password is
// percent-encoded. The previous string-formatted version broke on any password
// containing @ / ? # or :, which is not exotic — Neon and Railway both generate
// passwords from an alphabet that can include them, and the failure mode is a
// confusing "parse pool config" error naming a host that is only a fragment of
// the real one.
func DSN(cfg *config.Config) string {
	return buildDSN(cfg, cfg.Database.Host)
}

// MigrationDSN renders the connection string used for migrations only.
//
// It differs from DSN in one way, and the difference matters: it prefers a
// direct (unpooled) endpoint.
//
// golang-migrate serialises concurrent migrators with pg_advisory_lock, which
// is scoped to a *session*. Neon's pooled endpoint is PgBouncer in transaction
// mode, where a session does not map 1:1 onto a backend connection — the lock
// can be taken on one connection and the release issued on another. The result
// is a migration that hangs, or that fails partway and leaves schema_migrations
// flagged dirty, which then blocks every future migration until someone forces
// the version by hand.
//
// Neon's convention is that the pooled host is the direct host with "-pooler"
// appended to the endpoint id, so the direct host is recoverable by removing
// it. Only *.neon.tech hosts are touched; anything else is passed through
// untouched, since the suffix has no meaning elsewhere.
//
// MIGRATION_DATABASE_URL overrides all of this when set, for a provider whose
// pooled and direct endpoints are unrelated hostnames.
func MigrationDSN(cfg *config.Config) string {
	if override := strings.TrimSpace(cfg.Database.MigrationURL); override != "" {
		return override
	}
	return buildDSN(cfg, directHost(cfg.Database.Host))
}

// directHost strips Neon's "-pooler" marker from an endpoint hostname.
// ep-aged-darkness-1234-pooler.c-3.region.aws.neon.tech
//   → ep-aged-darkness-1234.c-3.region.aws.neon.tech
func directHost(host string) string {
	if !strings.HasSuffix(host, ".neon.tech") {
		return host
	}
	// Only the first label carries the marker; replacing globally could corrupt
	// a database name or region that legitimately contains the string.
	first, rest, found := strings.Cut(host, ".")
	if !found {
		return host
	}
	return strings.TrimSuffix(first, "-pooler") + "." + rest
}

func buildDSN(cfg *config.Config, host string) string {
	q := url.Values{}
	q.Set("sslmode", cfg.Database.SSLMode)

	u := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.Database.User, cfg.Database.Password),
		Host:     net.JoinHostPort(host, cfg.Database.Port),
		Path:     "/" + cfg.Database.Name,
		RawQuery: q.Encode(),
	}
	return u.String()
}
