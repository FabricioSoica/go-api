package main

import (
	"go-api/config"
	"go-api/routes"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRouter()
	r.Run(":8080")
}
