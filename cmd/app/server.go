package main

import (
	"github.com/gin-gonic/gin"
	"app/pkg/config"
	"app/pkg/routes"
	"log"
)

func main() {
	log.Println("Start serving")
	// Database
	config.Connect()
	// Init Router
	router := gin.Default()
	// Route Handlers / Endpoints
	routes.Routes(router)
	log.Fatal(router.Run(":8080"))
}
