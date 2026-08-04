package config

type Config struct {
	App AppConfig
	Server      ServerConfig
	Database    DatabaseConfig
	JWT         JWTConfig
	Cloudinary  CloudinaryConfig
}

type AppConfig struct{
	Name string
	Env string
}

type ServerConfig struct{
	Port        string
	CORSOrigins []string
}

type DatabaseConfig struct{
	Host 		string
	Port 		string
	User 		string
	Password 	string
	Name 		string
	SSLMode 	string

	// MigrationsPath is where the .sql migration files live. Defaults to
	// "migrations", which resolves correctly both from the repo root in
	// development and from / in the container, where the Dockerfile copies them
	// to /migrations.
	MigrationsPath string
	// MigrationURL, when set, is used verbatim for migrations instead of the
	// application connection. Only needed for a provider whose direct endpoint
	// is not derivable from the pooled one — see database.MigrationDSN.
	MigrationURL string
}

type JWTConfig struct{
	Secret string
}

type CloudinaryConfig struct {
	CloudName string
	APIKey    string
	APISecret string
}