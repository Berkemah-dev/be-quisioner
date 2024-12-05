package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// InitConfig memuat variabel lingkungan dari file .env
func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default environment variables")
	}
}

// GetEnv mengambil nilai variabel lingkungan berdasarkan nama
func GetEnv(envName string) string {
	value := os.Getenv(envName)
	if value == "" {
		log.Printf("Warning: Environment variable %s is not set", envName)
	}
	return value
}
