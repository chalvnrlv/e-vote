package main

import (
	"e-vote/config"
	"e-vote/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
