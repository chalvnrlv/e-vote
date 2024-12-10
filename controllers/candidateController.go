// controllers/candidateController.go

package controllers

import (
	"e-vote/config"
	"e-vote/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCandidate - creates a new candidate
func CreateCandidate(c *gin.Context) {
	var input models.Candidate

	// Bind JSON input to candidate struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if the Division ID exists in the Division table
	var division models.Division
	if err := config.GetDB().First(&division, input.DivisionID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Division not found"})
		return
	}

	// Set the Division reference to ensure GORM associates it correctly
	input.Division = division

	// Create the candidate
	if err := config.GetDB().Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create candidate"})
		return
	}

	// Return the created candidate with associated division
	c.JSON(http.StatusOK, gin.H{
		"message":   "Candidate created successfully",
		"candidate": input,
	})
}

// GetCandidates - fetches all candidates
func GetCandidates(c *gin.Context) {
	var candidatesWithDivisions []struct {
		ID         uint   `json:"id"`
		Name       string `json:"name"`
		DivisionID uint   `json:"division_id"` // This must match the SQL column alias
		Division   string `json:"division"`
	}

	// Using raw SQL to perform the JOIN and ensure correct aliases
	err := config.GetDB().Raw(`
        SELECT 
            candidates.ID, 
            candidates.Name, 
            candidates.Division_ID AS division_id, -- Ensure this alias matches the struct field
            divisions.Division
        FROM candidates
        LEFT JOIN divisions ON candidates.Division_ID = divisions.ID
    `).Scan(&candidatesWithDivisions).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch candidates"})
		return
	}

	// Return the fetched candidates with division data
	c.JSON(http.StatusOK, gin.H{"candidates": candidatesWithDivisions})
}

// UpdateCandidate - updates a candidate's information
func UpdateCandidate(c *gin.Context) {
	candidateID := c.Param("id")
	var candidateRequest models.Candidate

	// Bind the incoming JSON request to the candidate struct
	if err := c.ShouldBindJSON(&candidateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if the Division ID exists in the Division table if it's provided
	if candidateRequest.DivisionID != 0 {
		var division models.Division
		if err := config.GetDB().First(&division, candidateRequest.DivisionID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Division not found"})
			return
		}
	}

	// Update candidate in the database
	if err := config.GetDB().Where("id = ?", candidateID).Updates(&candidateRequest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update candidate"})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Candidate updated successfully"})
}

// DeleteCandidate - deletes a candidate
func DeleteCandidate(c *gin.Context) {
	candidateID := c.Param("id")

	// Delete the candidate from the database
	if err := config.GetDB().Where("id = ?", candidateID).Delete(&models.Candidate{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete candidate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Candidate deleted successfully"})
}
