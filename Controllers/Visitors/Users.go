package Visitors

import (
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"starter-golang/Application"
	"starter-golang/Models"
)

func Register(c *gin.Context) {
	r := Application.NewRequestWithAuth(c)
	var user Models.User
	r.Context.ShouldBindJSON(&user)

	// Validate Request
	err := validation.Errors{
		"name":     validation.Validate(user.Username, validation.Required.Error(gotrans.T("required")), validation.Length(5, 50).Error(gotrans.T("min_max"))),
		"email":    validation.Validate(user.Email, validation.Required.Error(gotrans.T("required")), is.Email.Error(gotrans.T("email")), validation.Length(5, 50).Error(gotrans.T("min_max"))),
		"password": validation.Validate(user.Password, validation.Required, validation.Length(5, 50).Error(gotrans.T("min_max"))),
	}.Filter()

	if err != nil {
		r.BadRequest(err)
		return
	}
	r.Ok(user)
}
