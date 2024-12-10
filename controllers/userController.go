package controllers

import (
	"e-vote/config"
	"e-vote/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser - creates a new user
func CreateUser(c *gin.Context) {
	var userRequest models.User

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Save the user to the database
	if err := config.GetDB().Create(&userRequest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": userRequest})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := config.GetDB().Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	var userRequest models.User

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Update user in the database
	if err := config.GetDB().Where("id = ?", userID).Updates(&userRequest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	// Delete the user from the database
	if err := config.GetDB().Where("id = ?", userID).Delete(&models.User{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
