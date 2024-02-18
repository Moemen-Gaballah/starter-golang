package Routes

import "starter-golang/Application"

type RouterApp struct {
	*Application.Application
}

func (app RouterApp) Routing() {
	app.visitorsRoutes()
	app.authRoutes()
	app.adminsRoutes()
}
