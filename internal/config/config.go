package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr string
	DBUrl string
}

func Load() *Config {
	godotenv.Load()
	addr := os.Getenv("ADDR")
	if addr == "" {
        addr = ":5000"
    }

	dbUrl := os.Getenv("DB_URL")

	return &Config{
		Addr: addr,
		DBUrl: dbUrl,
	}
}