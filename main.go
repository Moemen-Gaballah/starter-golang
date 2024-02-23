package main

import (
	"starter-golang/Application"
	"starter-golang/Models"
	"starter-golang/Routes"
)

func main() {
	app := Application.NewApp()

	// migrate project
	app.DB.AutoMigrate(&Models.User{})

	// close app connection
	Application.CloseConnection(app)

	// start routing
	routeApp := Routes.RouterApp{app}
	routeApp.Routing()

	app.Gin.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
