package routes

import (
	"e-vote/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS middleware
	r.Use(corsMiddleware())

	// User routes
	r.OPTIONS("/login", handleOptionsRequest)
	r.POST("/login", controllers.Login)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Candidate routes
	r.POST("/candidates", controllers.CreateCandidate)
	r.GET("/candidates", controllers.GetCandidates)
	r.PUT("/candidates/:id", controllers.UpdateCandidate)
	r.DELETE("/candidates/:id", controllers.DeleteCandidate)

	// Divisions routes
	r.GET("/divisions", controllers.GetDivisions)
	r.GET("/divisions/:id/candidates", controllers.GetCandidatesByDivision)
	r.POST("/divisions", controllers.CreateDivision)
	r.PUT("/divisions/:id", controllers.UpdateDivision)
	r.DELETE("/divisions/:id", controllers.DeleteDivision)

	// Vote route
	r.GET("/vote/:users.id", controllers.GetUserVotes)
	r.POST("/vote/div-:divisions.id/:users.id", controllers.Vote)

	r.GET("/results/:division_id", controllers.GetDivisionResults)

	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, Accept")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

// COOOOOOOOOORS
func handleOptionsRequest(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Status(http.StatusOK)
}
