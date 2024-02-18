package Routes

import "starter-golang/Controllers/Visitors"

func (app RouterApp) visitorsRoutes() {
	app.Gin.GET("/create-user", Visitors.CreateUser)
}
