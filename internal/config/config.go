package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr string
}

func Load() *Config {
	godotenv.Load()
	addr := os.Getenv("ADDR")
	if addr == "" {
        addr = ":5000"
    }

	return &Config{
		Addr: addr,
	}
}