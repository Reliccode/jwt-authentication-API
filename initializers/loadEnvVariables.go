package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnvVariables loads environment variables from a .env file.
// It uses the godotenv package to load the variables.
// Fatal errror will be logged if loading fails. 
func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
