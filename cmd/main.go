package main

import (
	"e-vote/config"
	"e-vote/routes"
)

func main() {
	config.InitDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
