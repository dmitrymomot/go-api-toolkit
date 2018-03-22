package validator

import (
	validator "gopkg.in/go-playground/validator.v9"
)

func phoneNumber(fl validator.FieldLevel) bool {
	return IsPhoneNumber(fl.Field().String())
}
