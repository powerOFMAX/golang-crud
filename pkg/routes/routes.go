package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"app/pkg/controllers"
)

func Routes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", welcome)
		v1.GET("/messages", controllers.GetAllMessages)
		v1.POST("/message", controllers.CreateMessage)
		v1.GET("/message/:messageId", controllers.GetSingleMessage)
		v1.PUT("/message/:messageId", controllers.EditMessage)
		v1.DELETE("/message/:messageId", controllers.DeleteMessage)
	}
	router.Run()
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}