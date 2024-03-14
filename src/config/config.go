package config

import "os"

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBHost     string
	DB         string
}

func NewConfig() *Config {
	return &Config{
		Port:       os.Getenv("PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DB:         os.Getenv("DB"),
	}
}
