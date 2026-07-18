package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Load() (*Config, error){
	_ = godotenv.Load()

	dbHost, err := getEnv("DB_HOST")
	if err != nil {
		return nil, err
	}

	dbPort, err := getEnv("DB_PORT")
	if err != nil {
		return nil, err
	}

	dbUser, err := getEnv("DB_USER")
	if err != nil {
		return nil, err
	}

	dbPassword, err := getEnv("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	dbName, err := getEnv("DB_NAME")
	if err != nil {
		return nil, err
	}

	jwtSecret, err := getEnv("JWT_SECRET")
	if err != nil {
		return nil, err
	}

	appName := getEnvOrDefault("APP_NAME", "BBS Backend")

	appEnv := getEnvOrDefault("APP_ENV", "development")

	serverPort := getEnvOrDefault("SERVER_PORT", "8080")

	sslMode := getEnvOrDefault("SSL_MODE", "disable")

	cfg := &Config{
		App: AppConfig{
			Name: appName,
			Env: appEnv,
		},
		Server: ServerConfig{
			Port: serverPort,
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
	}

	return cfg, nil
}

func getEnv(key string) (string, error){
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("%s is required", key)
	}

	return value, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == ""{
		return defaultValue
	}

	return value
}
