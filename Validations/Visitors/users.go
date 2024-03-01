package Visitors

import (
	"github.com/bykovme/gotrans"
	validation "github.com/go-ozzo/ozzo-validation"
	"starter-golang/Models"
	"starter-golang/Validations"
)

func RegisterValidation(user Models.User) validation.Errors {
	return validation.Errors{
		"name":     validation.Validate(user.Username, Validations.RequiredRule(), validation.Length(5, 50).Error(gotrans.T("min_max"))),
		"email":    validation.Validate(user.Email, Validations.RequiredRule(), Validations.IsEmailRule(), Validations.MinMaxRule()),
		"password": validation.Validate(user.Password, Validations.RequiredRule(), Validations.MinMaxRule()),
	}
}

func LoginValidation(user Models.User) validation.Errors {
	return validation.Errors{
		"email":    validation.Validate(user.Email, Validations.RequiredRule(), Validations.IsEmailRule(), Validations.MinMaxRule()),
		"password": validation.Validate(user.Password, Validations.RequiredRule(), Validations.MinMaxRule()),
	}
}
