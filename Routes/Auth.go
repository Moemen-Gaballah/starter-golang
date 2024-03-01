package Routes

import "starter-golang/Controllers/Auth"

func (app RouterApp) authRoutes() {
	app.Gin.GET("/me", Auth.Me)
}
