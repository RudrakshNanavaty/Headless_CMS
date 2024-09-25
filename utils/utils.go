package utils

// create an object and get all env variables and export
import (
	"github.com/joho/godotenv"
	"os"
)

// LoadEnv function to load environment variables
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}

// GetEnv function to get environment variables
func GetEnv(key string) string {
	return os.Getenv(key)
}
