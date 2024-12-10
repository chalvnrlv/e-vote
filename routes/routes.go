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

	return r
}
