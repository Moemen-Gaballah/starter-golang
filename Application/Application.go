package App

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Application struct {
	Gin        *gin.Engine
	DB         *gorm.DB
	Connection *sql.DB
}

func (app *Application) Share() {

}

func App() func() Application {
	return func() Application {
		var application Application
		application.Gin = gin.Default()
		connectToDatabase(&application)
		return application
	}
}
