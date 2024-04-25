package configs

import (
	"os"
)

// Config holds application configuration settings
type DBstr struct {
	// Database settings
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// NewDBConfig initializes a new DBConfig instance
func DBConfig() *DBstr {
	return &DBstr{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}
