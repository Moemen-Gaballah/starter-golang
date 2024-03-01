package Seeders

import (
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
	"starter-golang/Models"
)

func seedUsers(DB *gorm.DB) {
	DB.Create(&Models.User{
		Username: gofakeit.Name(),
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, true, true, 8),
		Group:    "admin",
		Token:    gofakeit.Email(),
	})
}
