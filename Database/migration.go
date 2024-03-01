package Database

import (
	"gorm.io/gorm"
	"starter-golang/Models"
)

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(
		&Models.User{},
	)
}
