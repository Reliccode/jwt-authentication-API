package models

import "gorm.io/gorm"

// User represents the user model in the application.
// Structure definitin of user data stored in the database.
type User struct {
	gorm.Model // gorm.Model is embedded to provide common fields like ID, CreatedAt, UpdatedAt, and DeletedAt
	Email           string `gorm:"unique"` // Email field represents the email address of the user and is marked as unique in the database
	Password   string // Password field stores the hashed password of the user
}