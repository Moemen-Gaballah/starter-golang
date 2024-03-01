package Routes

import "starter-golang/Controllers/Visitors"

func (app RouterApp) visitorsRoutes() {
	app.Gin.POST("register", Visitors.Register)
}
