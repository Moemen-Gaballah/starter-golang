package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	app := app()
	application := app()
	application.Gin.GET("/ping", func(c *gin.Context) {

		r := newRequest(c)
		r.Connection.Close()

		//r.Response(200, gin.H{
		//	"message": "Hello Every One",
		//})
		r.Ok(gin.H{
			"message": "Hello Every One",
		})

		//request.Context.JSON(http.StatusOK, gin.H{
		//	"message": "pong",
		//})
	})

	application.Gin.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
