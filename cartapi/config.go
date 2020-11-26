package cartapi

import (
	"github.com/joho/godotenv"
	"os"
)

var Config config

type config struct {
	Server      ServerConfig
	Database    DatabaseConfig
	Environment string
}

type ServerConfig struct {
	Addr string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Schema   string
	SslMode  string
}

func LoadConfig() {
	env := os.Getenv("CARTAPI_ENV")
	if "" == env {
		env = "development"
	}

	loadFiles(env)
	Config = config{
		Server: ServerConfig{
			Addr: os.Getenv("SERVER_ADDR"),
		},
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
			Schema:   os.Getenv("DB_SCHEMA"),
			SslMode:  os.Getenv("DB_SSL_MODE"),
		},
		Environment: env,
	}
}

func loadFiles(env string) {
	_ = godotenv.Load(".env." + env + ".local")
	if "test" != env {
		_ = godotenv.Load(".env.local")
	}

	if err := godotenv.Load(".env." + env); err != nil {
		panic("failed to load env configuration file")
	}
	if err := godotenv.Load(); err != nil {
		panic("failed to load base configuration file")
	}
}
