package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	app := app()
	application := app()
	application.Gin.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	application.Gin.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
