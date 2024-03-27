package main

// Package main is the entry point of the application.
// It initializes the application by loading environment variables, connecting to the database,
// and synchronizing the database schema with defined models.
// Then, it sets up HTTP routes for user signup, login, and validation,
// utilizing controllers to handle business logic and middleware for authentication.
// Finally, it starts the server to listen and serve incoming HTTP requests.

import (
	"github.com/Reliccode/jwt-authentication-API/controllers"
	"github.com/Reliccode/jwt-authentication-API/initializers"
	"github.com/Reliccode/jwt-authentication-API/middleware"
	"github.com/gin-gonic/gin"
)

// init initializes the application by loading environment variables, connecting to the database,
// and synchronizing database schema with defined models.
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	// Initialize a new Gin router
	r := gin.Default()

	// Define routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// Run the server
	r.Run()
}
