package database

import (
	"strings"
	"testing"

	"github.com/dedehudianto12/bbs-backend/config"
)

func TestDirectHost(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "neon pooled endpoint loses the marker",
			in:   "ep-aged-darkness-azrsrgk7-pooler.c-3.ap-southeast-1.aws.neon.tech",
			want: "ep-aged-darkness-azrsrgk7.c-3.ap-southeast-1.aws.neon.tech",
		},
		{
			name: "neon direct endpoint is already correct",
			in:   "ep-aged-darkness-azrsrgk7.c-3.ap-southeast-1.aws.neon.tech",
			want: "ep-aged-darkness-azrsrgk7.c-3.ap-southeast-1.aws.neon.tech",
		},
		{
			name: "non-neon host is untouched even when it contains the marker",
			in:   "my-pooler.example.com",
			want: "my-pooler.example.com",
		},
		{
			name: "localhost is untouched",
			in:   "localhost",
			want: "localhost",
		},
		{
			// The marker must only be stripped from the endpoint label. A region
			// or account label ending in -pooler must survive.
			name: "only the first label is rewritten",
			in:   "ep-x-pooler.something-pooler.aws.neon.tech",
			want: "ep-x.something-pooler.aws.neon.tech",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := directHost(tc.in); got != tc.want {
				t.Errorf("directHost(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

// A password containing URL-significant characters must survive into the DSN
// intact. The fmt.Sprintf version this replaced would have produced a string
// that parsed as a completely different host.
func TestBuildDSNEscapesPassword(t *testing.T) {
	cfg := &config.Config{}
	cfg.Database.User = "neondb_owner"
	cfg.Database.Password = "p@ss:w/rd?x#y"
	cfg.Database.Host = "db.example.com"
	cfg.Database.Port = "5432"
	cfg.Database.Name = "neondb"
	cfg.Database.SSLMode = "require"

	got := buildDSN(cfg, cfg.Database.Host)

	if strings.Contains(got, "p@ss:w/rd?x#y") {
		t.Fatalf("password was not escaped: %q", got)
	}
	if !strings.Contains(got, "@db.example.com:5432/neondb") {
		t.Errorf("host/database mangled: %q", got)
	}
	if !strings.Contains(got, "sslmode=require") {
		t.Errorf("sslmode missing: %q", got)
	}
}

// MigrationDSN must prefer the direct endpoint, and an explicit override must
// win over the derivation.
func TestMigrationDSN(t *testing.T) {
	cfg := &config.Config{}
	cfg.Database.User = "u"
	cfg.Database.Password = "p"
	cfg.Database.Host = "ep-abc-pooler.c-3.ap-southeast-1.aws.neon.tech"
	cfg.Database.Port = "5432"
	cfg.Database.Name = "neondb"
	cfg.Database.SSLMode = "require"

	got := MigrationDSN(cfg)
	if strings.Contains(got, "-pooler") {
		t.Errorf("migration DSN still points at the pooled endpoint: %q", got)
	}

	// The application pool, by contrast, should keep using the pooler.
	if app := DSN(cfg); !strings.Contains(app, "-pooler") {
		t.Errorf("application DSN should keep the pooled endpoint: %q", app)
	}

	cfg.Database.MigrationURL = "postgres://override/db"
	if got := MigrationDSN(cfg); got != "postgres://override/db" {
		t.Errorf("MIGRATION_DATABASE_URL override ignored: %q", got)
	}
}
