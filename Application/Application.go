package Application

import (
	"database/sql"
	"fmt"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"starter-golang/Database"
)

type Application struct {
	Gin        *gin.Engine
	DB         *gorm.DB
	Connection *sql.DB
}

func (app *Application) Share() {

}

func App() func() *Application {
	return func() *Application {
		var application Application
		application.Gin = gin.Default()
		connectToDatabase(&application)

		err := gotrans.InitLocales("./public/lang") //  Path to the folder with localization files
		if err != nil {
			fmt.Println("Error on load trans files", err.Error())
		}
		return &application
	}
}

// init new request closure
func NewApp() *Application {
	app := App()
	return app()
}

func (app *Application) Migrate() {
	Database.Migrate(app.DB)
}
