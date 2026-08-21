package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	DSN string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config.")
	}

	return &Config{
		Db: DbConfig{
			DSN: os.Getenv("DB_DSN"),
		},
	}
}
