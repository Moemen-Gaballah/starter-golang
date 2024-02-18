package Application

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

// init new request closure
func NewApp() Application {
	app := App()
	application := app()

	return application
}
