package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppUrl      string
	Port        string
	DatabaseUrl string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Could not load .env file", err)
	}

	appUrl := os.Getenv("APP_URL")
	if appUrl == "" {
		appUrl = "development"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	return Config{
		AppUrl:      appUrl,
		Port:        port,
		DatabaseUrl: dbUrl,
	}
}
