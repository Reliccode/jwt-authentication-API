package initializers

// initializers contains functions responsible for initialization of various
// components of the application, such as database connections.

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB holds the connection to the database.
var DB *gorm.DB

// ConnectToDb initializes a connection to the database.
// It uses the value of the "DB" environment variable to determine the database connection string.
// If the connection fails, it panics.
func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}
}