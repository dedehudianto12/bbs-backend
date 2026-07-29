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
}

type JWTConfig struct{
	Secret string
}

type CloudinaryConfig struct {
	CloudName string
	APIKey    string
	APISecret string
}