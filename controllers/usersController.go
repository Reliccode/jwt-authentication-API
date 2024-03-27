package controllers

// The controllers package contains functions responsible for handling HTTP requests and
// implementing the business logic of the application related to user authentication.

import (
	"net/http"
	"os"
	"time"

	"github.com/Reliccode/jwt-authentication-API/initializers"
	"github.com/Reliccode/jwt-authentication-API/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
)

// Signup handles the user signup process.
// It reads email and password from the request body, hashes the password,
// creates a new user, and responds with success or failure.
func Signup(c *gin.Context) {
	var body struct {
		Email         string
		Password string
	}

	if c.Bind(&body) != nil {
		// If failed to bind request body, respond with bad request
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		// If failed to hash password, respond with bad request
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		// If failed to create user, respond with bad request
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{})
}

// Login handles the user login process.
// It retrieves email and password from the request body,
// validates them against the stored user credentials,
// generates a JWT token upon successful login, sets it in a cookie,
// and responds with success or failure.
func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		// If failed to bind request body, respond with bad request
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		// If user not found, respond with bad request
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Compare sent password with saved user password hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		// If password comparison fails, respond with bad request
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		// If token creation fails, respond with bad request
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Set token in a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	// Respond with success
	c.JSON(http.StatusOK, gin.H{})
}

// Validate validates the user.
// It expects a user object in the context and responds with the user's details.
func Validate(c *gin.Context) {
	user, _ := c.Get("user") // Assuming "user" is set during authentication

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}