package config

type Config struct {
	App AppConfig
	Server ServerConfig
	Database DatabaseConfig
	JWT JWTConfig
}

type AppConfig struct{
	Name string
	Env string
}

type ServerConfig struct{
	Port string
}

type DatabaseConfig struct{
	Host 		string
	Port 		string
	User 		string
	Password 	string
	Name 		string
	SSLMode 	string
}

type JWTConfig struct{
	Secret string
}