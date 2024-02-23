package Routes

import "starter-golang/Controllers/Auth"

func (app RouterApp) authRoutes() {
	app.Gin.GET("/create-user", Auth.CreateUser)
}
