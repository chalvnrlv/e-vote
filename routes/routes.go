// routes/routes.go

package routes

import (
	"e-vote/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User routes
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
	r.GET("/divisions/:id", controllers.GetCandidatesByDivision)
	r.POST("/divisions", controllers.CreateDivision)
	r.PUT("/divisions/:id", controllers.UpdateDivision)
	r.DELETE("/divisions/:id", controllers.DeleteDivision)

	// Vote route
	r.GET("/vote/:users.id", controllers.GetUserVotes)
	r.POST("/vote/div-:divisions.id/:users.id", controllers.Vote)

	r.GET("/results/:division_id", controllers.GetDivisionResults)

	return r
}
