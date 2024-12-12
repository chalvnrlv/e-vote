package controllers

import (
	"e-vote/config"
	"e-vote/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Vote(c *gin.Context) {
	// Extract division ID and user ID from URL parameters
	divisionIDStr := c.Param("divisions.id")
	userIDStr := c.Param("users.id")

	// Convert divisionID and userID from string to uint
	divisionID, err := strconv.ParseUint(divisionIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid division ID"})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Parse Candidate_ID from the request body
	var voteData struct {
		CandidateID uint `json:"Candidate_ID"`
	}
	if err := c.ShouldBindJSON(&voteData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if the division exists
	var division models.Division
	if err := config.DB.First(&division, "id = ?", divisionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Division not found"})
		return
	}

	// Check if the candidate exists in this division
	var candidate models.Candidate
	if err := config.DB.First(&candidate, "id = ? AND division_id = ?", voteData.CandidateID, divisionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found in this division"})
		return
	}

	// Check if the user has already voted in this division
	var existingVote models.UserCandidate
	if err := config.DB.Where("user_id = ? AND candidate_id IN (?)", userID, config.DB.Model(&models.Candidate{}).Where("division_id = ?", divisionID).Select("id")).First(&existingVote).Error; err == nil {
		// If there's an existing vote, the user cannot vote again in the same division
		c.JSON(http.StatusConflict, gin.H{"error": "You have already voted in this division"})
		return
	}

	// Create a new vote in the user_candidates table
	vote := models.UserCandidate{
		UserID:      uint(userID), // Convert uint from userID
		CandidateID: voteData.CandidateID,
	}
	if err := config.DB.Create(&vote).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error casting vote"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{
		"message":      "Vote cast successfully",
		"user_id":      userID,
		"candidate_id": voteData.CandidateID,
	})
}
func GetUserVotes(c *gin.Context) {
	userIDStr := c.Param("users.id")

	// Convert userID from string to uint
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Fetch all divisions
	var divisions []models.Division
	if err := config.DB.Find(&divisions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching divisions"})
		return
	}

	// Fetch the user's votes
	var userVotes []struct {
		DivisionID uint   `json:"division_id"`
		Division   string `json:"division"`
		Candidate  string `json:"candidate"`
	}
	query := `
		SELECT d.id AS division_id, d.division, c.name AS candidate
		FROM user_candidates uc
		JOIN candidates c ON uc.candidate_id = c.id
		JOIN divisions d ON c.division_id = d.id
		WHERE uc.user_id = ?
	`
	if err := config.DB.Raw(query, userID).Scan(&userVotes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user votes"})
		return
	}

	// Map divisions the user has voted for
	votedDivisions := map[uint]bool{}
	for _, vote := range userVotes {
		votedDivisions[vote.DivisionID] = true
	}

	// Determine if the user has voted in all divisions
	missingDivisions := []string{}
	for _, division := range divisions {
		if !votedDivisions[division.ID] {
			missingDivisions = append(missingDivisions, division.Division)
		}
	}

	// Build the response
	if len(userVotes) == 0 {
		// User has not voted at all
		c.JSON(http.StatusOK, gin.H{
			"message": "VOTE NOW",
			"votes":   userVotes,
		})
		return
	}

	if len(missingDivisions) == 0 {
		// User has voted in all divisions
		c.JSON(http.StatusOK, gin.H{
			"message": "You have voted in all divisions",
			"votes":   userVotes,
		})
		return
	}

	// User has voted in some divisions but not all
	missingMessage := fmt.Sprintf("You haven't voted yet in %v", missingDivisions)
	c.JSON(http.StatusOK, gin.H{
		"message": missingMessage,
		"votes":   userVotes,
	})
}

func GetDivisionResults(c *gin.Context) {
	// Get division_id from the URL parameter
	divisionIDStr := c.Param("division_id")
	divisionID, err := strconv.Atoi(divisionIDStr)
	if err != nil || divisionID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid division ID"})
		return
	}

	// Get the division details
	var division models.Division
	if err := config.DB.Where("id = ?", divisionID).First(&division).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Division not found"})
		return
	}

	// Get all candidates in the division and their vote counts
	var candidateResults []struct {
		CandidateName string `json:"candidate_name"`
		VoteCount     int    `json:"vote_count"`
	}

	query := `
		SELECT c.name AS candidate_name, COUNT(uc.user_id) AS vote_count
		FROM candidates c
		LEFT JOIN user_candidates uc ON c.id = uc.candidate_id
		WHERE c.division_id = ?
		GROUP BY c.id
		ORDER BY vote_count DESC
	`

	if err := config.DB.Raw(query, divisionID).Scan(&candidateResults).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching division results"})
		return
	}

	// Return the division results
	c.JSON(http.StatusOK, gin.H{
		"division": division.Division,
		"results":  candidateResults,
	})
}
