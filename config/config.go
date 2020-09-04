package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port string
	Password string
	DB *DBConfig
}

type DBConfig struct {
	Username string
	Password string
	Host string
	Port string
	Name string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load .env file!")
	}

	return &Config{
		Port: os.Getenv("APP_PORT"),
		Password: os.Getenv("APP_PASSWORD"),
		DB: &DBConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Name: os.Getenv("DB_NAME"),
		},
	}
}