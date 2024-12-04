package routes

import (
	"e-vote/controllers"
	"e-vote/utils"

	"github.com/gin-gonic/gin"
)

// AdminMiddleware checks if the logged-in user is an admin (RoleID == 1)
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(400, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		// Parse the token
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Check if the user is an admin (RoleID == 1)
		if claims.Role != "1" {
			c.JSON(403, gin.H{"error": "You do not have permission to access this route"})
			c.Abort()
			return
		}

		// Proceed to the next handler
		c.Next()
	}
}

func SetupRoutes(router *gin.Engine) {
	// Public route for login
	router.POST("/login", controllers.Login)

	// Protected routes (requires JWT authentication)
	authorized := router.Group("/")
	authorized.Use(utils.JWTAuthMiddleware()) // Use JWT middleware to authenticate

	// Voting route (all authenticated users can vote)
	authorized.POST("/vote", controllers.Vote)

	// Admin-only routes (use AdminMiddleware to check if the user is an admin)
	authorizedAdmin := router.Group("/")
	authorizedAdmin.Use(utils.JWTAuthMiddleware(), AdminMiddleware()) // JWT + Admin check
	{
		// Admin can create candidates (CRUD)
		authorizedAdmin.POST("/candidate", controllers.CreateCandidate)

		// Admin can add users (for testing purposes or user management)
		authorizedAdmin.POST("/create-user", controllers.CreateUser)
	}
}
