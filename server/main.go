package main

import (
	//"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	// Ping test
	server.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})


	// Listen and Server in 0.0.0.0:8080
	server.Run(":8080")

}
