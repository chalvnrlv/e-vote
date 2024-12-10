package controllers

import (
	"e-vote/config" // Import config to access DB
	"e-vote/models"
	"net/http"
	"strconv"

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

func GetCandidatesByDivision(c *gin.Context) {
	divisionID := c.Param("id")

	// Ensure divisionID is an integer
	divisionIDInt, err := strconv.Atoi(divisionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid division ID"})
		return
	}

	var candidates []models.Candidate

	// Query candidates by division ID, preload the Division and select only necessary fields
	if err := config.DB.Preload("Division").Select("id", "name", "image").Where("division_id = ?", divisionIDInt).Find(&candidates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching candidates for division"})
		return
	}

	// If no candidates found
	if len(candidates) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No candidates found for this division"})
		return
	}

	// Return only id, name, and image in the response
	result := []gin.H{}
	for _, candidate := range candidates {
		result = append(result, gin.H{
			"id":    candidate.ID,
			"name":  candidate.Name,
			"image": candidate.Image, // Assuming image is in byte array
		})
	}

	c.JSON(http.StatusOK, result)
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

func UpdateDivision(c *gin.Context) {
	divisionID := c.Param("id")
	var division models.Division

	// Check if the division exists
	if err := config.DB.First(&division, divisionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Division not found"})
		return
	}

	// Bind the new data for the division
	if err := c.ShouldBindJSON(&division); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the division
	if err := config.DB.Save(&division).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating division"})
		return
	}

	c.JSON(http.StatusOK, division)
}

func DeleteDivision(c *gin.Context) {
	divisionID := c.Param("id")
	var division models.Division

	// Check if the division exists
	if err := config.DB.First(&division, divisionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Division not found"})
		return
	}

	// Delete the candidates associated with this division
	if err := config.DB.Where("division_id = ?", divisionID).Delete(&models.Candidate{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting candidates for division"})
		return
	}

	// Delete the division
	if err := config.DB.Delete(&division).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting division"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Division and associated candidates deleted successfully"})
}
