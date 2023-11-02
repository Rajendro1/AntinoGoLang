package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
}

func DbConfig() DBConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	c := DBConfig{
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
	}
	return c
}

var (
	DB_NAME = "blogs"
)
