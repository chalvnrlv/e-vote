package controllers

import (
	"e-vote/config" // Import config to access DB
	"e-vote/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDivisions(c *gin.Context) {
	var divisions []models.Division
	if err := config.DB.Find(&divisions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching divisions"})
		return
	}
	c.JSON(http.StatusOK, divisions)
}

func CreateDivision(c *gin.Context) {
	var division models.Division
	if err := c.ShouldBindJSON(&division); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&division).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating division"})
		return
	}

	c.JSON(http.StatusCreated, division)
}
