package config

import (
	"log"
	"os"
)

type Config struct {
	Port        string
	DatabaseUrl string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	return Config{
		Port:        port,
		DatabaseUrl: dbUrl,
	}
}
