package initializers

import (
	"github.com/Reliccode/jwt-authentication-API/models"
)

// SyncDatabase performs automatic migration of database tables.
// It ensures that the database schema matches the structure of the defined models.
func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}