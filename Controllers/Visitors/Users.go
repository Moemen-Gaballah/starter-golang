package Visitors

import (
	"github.com/gin-gonic/gin"
	"starter-golang/Application"
	"starter-golang/Models"
	"starter-golang/Validations/Visitors"
)

func Register(c *gin.Context) {
	r := Application.NewRequest(c)
	var user Models.User
	r.Context.ShouldBindJSON(&user)

	// Validate Request
	r.ValidateRequest(Visitors.RegisterValidation(user))

	if r.Fails() {
		return
	}
	user.Token = user.Email
	user.Group = "user"
	r.DB.Create(&user)
	r.Created(user)
}
