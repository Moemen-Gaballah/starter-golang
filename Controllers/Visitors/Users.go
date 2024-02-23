package Visitors

import (
	"github.com/gin-gonic/gin"
	"starter-golang/Application"
	"starter-golang/Models"
)

func CreateUser(c *gin.Context) {

	r := Application.NewRequest(c).Auth()
	if !r.IsAdmin {
		r.NotAuth()
		return
	}

	user := Models.User{
		Username: "Moemen Gaballa",
		Email:    "moemen@gmail.com",
		Password: "123456",
	}

	r.DB.Create(&user)
	r.Created(user)
}
