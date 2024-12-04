package controllers

import (
	"e-vote/config"
	"e-vote/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCandidate(c *gin.Context) {
	var input models.Candidate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := config.GetDB().Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create candidate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Candidate created successfully"})
}
