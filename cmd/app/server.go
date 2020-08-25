package main

import (
	"app/pkg/config"
	"app/pkg/routes"
	"github.com/gin-gonic/gin"
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
