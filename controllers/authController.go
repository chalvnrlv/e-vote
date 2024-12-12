package controllers

import (
	"e-vote/config"
	"e-vote/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Login function to handle user authentication
func Login(c *gin.Context) {
	var input models.User

	// Bind the incoming JSON request to the input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Look for a user in the database with the given identity_number
	var user models.User
	if err := config.GetDB().Where("IdentityNumber = ?", input.IdentityNumber).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Check if the passwords match
	if user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check user role and redirect accordingly
	var redirectURL string
	if user.RoleID == 1 {
		// Admin role
		redirectURL = "/admin/dashboard"
	} else {
		// Regular user role
		redirectURL = "/vote/" + fmt.Sprintf("%d", user.ID)
	}

	// Return success response with user data and redirect URL
	c.JSON(http.StatusOK, gin.H{
		"message":     "Login successful",
		"user":        user,
		"redirectURL": redirectURL,
	})
}
