package controllers

import (
	"e-vote/config"
	"e-vote/models"
	"e-vote/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	if err := config.GetDB().Where("IdentityNumber = ?", input.IdentityNumber).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid identity number or password"})
		return
	}

	// Log the user data to verify it's being retrieved correctly
	fmt.Println("User retrieved:", user)

	// Temporarily remove bcrypt check for plain-text password comparison
	if user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid identity number or password"})
		return
	}

	// Generate the token if login is successful
	token, err := utils.GenerateToken(user.ID, user.Name, fmt.Sprint(user.RoleID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func CreateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// // Hash the password before saving it to the database
	// if err := input.HashPassword(); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
	// 	return
	// }

	// Save the user to the database
	if err := config.GetDB().Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
