package main

import "github.com/gin-gonic/gin"

type Application struct {
	Gin *gin.Engine
}

func app() func() Application {
	return func() Application {
		var application Application
		application.Gin = gin.Default()

		return application
	}
}
