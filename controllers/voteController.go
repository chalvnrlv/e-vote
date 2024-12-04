package controllers

import (
	"e-vote/config"
	"e-vote/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Vote(c *gin.Context) {
	userID := c.Param("user_id")
	var input struct {
		CandidateID uint `json:"candidate_id"`
		DivisionID  uint `json:"division_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var existingVote models.UserCandidate
	if err := config.GetDB().Where("user_id = ? AND candidate_id = ?", userID, input.CandidateID).First(&existingVote).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You already voted for this candidate"})
		return
	}

	vote := models.UserCandidate{
		UserID:      input.CandidateID,
		CandidateID: input.CandidateID,
	}

	if err := config.GetDB().Create(&vote).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit vote"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vote submitted"})
}
