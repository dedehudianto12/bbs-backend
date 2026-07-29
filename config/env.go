package config

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Load() (*Config, error){
	_ = godotenv.Load()

	// DATABASE_URL takes precedence when set (Railway / Neon style)
	var dbHost, dbPort, dbUser, dbPassword, dbName, sslMode string
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		var err error
		dbHost, dbPort, dbUser, dbPassword, dbName, sslMode, err = parseDatabaseURL(dbURL)
		if err != nil {
			return nil, fmt.Errorf("DATABASE_URL: %w", err)
		}
	} else {
		var err error
		dbHost, err = getEnv("DB_HOST")
		if err != nil {
			return nil, err
		}
		dbPort, err = getEnv("DB_PORT")
		if err != nil {
			return nil, err
		}
		dbUser, err = getEnv("DB_USER")
		if err != nil {
			return nil, err
		}
		dbPassword, err = getEnv("DB_PASSWORD")
		if err != nil {
			return nil, err
		}
		dbName, err = getEnv("DB_NAME")
		if err != nil {
			return nil, err
		}
		sslMode = getEnvOrDefault("SSL_MODE", "disable")
	}

	jwtSecret, err := getEnv("JWT_SECRET")
	if err != nil {
		return nil, err
	}

	appName := getEnvOrDefault("APP_NAME", "BBS Backend")

	appEnv := getEnvOrDefault("APP_ENV", "development")

	serverPort := getEnvOrDefault("PORT", getEnvOrDefault("SERVER_PORT", "8080"))

	corsOrigins := parseCORSOrigins(getEnvOrDefault("CORS_ORIGINS", "http://localhost:3000,http://127.0.0.1:3000"))

	cloudinaryCloudName := getEnvOrDefault("CLOUDINARY_CLOUD_NAME", "")
	cloudinaryAPIKey := getEnvOrDefault("CLOUDINARY_API_KEY", "")
	cloudinaryAPISecret := getEnvOrDefault("CLOUDINARY_API_SECRET", "")

	cfg := &Config{
		App: AppConfig{
			Name: appName,
			Env: appEnv,
		},
		Server: ServerConfig{
			Port:        serverPort,
			CORSOrigins: corsOrigins,
		},
		Database: DatabaseConfig{
			Host: dbHost,
			Port: dbPort,
			User: dbUser,
			Password: dbPassword,
			Name: dbName,
			SSLMode: sslMode,
		},
		JWT: JWTConfig{
			Secret: jwtSecret,
		},
		Cloudinary: CloudinaryConfig{
			CloudName: cloudinaryCloudName,
			APIKey:    cloudinaryAPIKey,
			APISecret: cloudinaryAPISecret,
		},
	}

	return cfg, nil
}

func parseCORSOrigins(raw string) []string {
	parts := strings.Split(raw, ",")
	origins := make([]string, 0, len(parts))
	for _, p := range parts {
		if t := strings.TrimSpace(p); t != "" {
			origins = append(origins, t)
		}
	}
	return origins
}

func getEnv(key string) (string, error){
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("%s is required", key)
	}

	return value, nil
}

// parseDatabaseURL parses a postgres:// URL into its components.
func parseDatabaseURL(raw string) (host, port, user, password, name, sslMode string, err error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", "", "", "", "", "", fmt.Errorf("invalid URL: %w", err)
	}
	host = u.Hostname()
	port = u.Port()
	if port == "" {
		port = "5432"
	}
	user = u.User.Username()
	password, _ = u.User.Password()
	name = strings.TrimPrefix(u.Path, "/")
	sslMode = u.Query().Get("sslmode")
	if sslMode == "" {
		sslMode = "require" // production default for cloud databases
	}
	return
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == ""{
		return defaultValue
	}

	return value
}
