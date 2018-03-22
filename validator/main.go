package validator

import (
	"errors"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

// Validator instance
var v = &struct {
	v   *validator.Validate
	uni *ut.UniversalTranslator
	t   ut.Translator
}{}

// Setup validator
func Setup() error {
	en := en.New()
	v.uni = ut.New(en.(locales.Translator), en.(locales.Translator))
	tr, found := v.uni.GetTranslator("en")
	if !found {
		return errors.New("Could not found translator")
	}
	v.t = tr
	v.v = validator.New()

	v.v.RegisterValidation("phone", phoneNumber)
	v.t.Add("phone", "Must be a valid phone number", false)

	en_translations.RegisterDefaultTranslations(v.v, v.t)

	return nil
}

// Struct validates structure
func Struct(s interface{}) error {
	if v == nil {
		if err := Setup(); err != nil {
			return err
		}
	}
	err := v.v.Struct(s)
	if err != nil {
		ve := err.(validator.ValidationErrors)
		return &Errors{Errors: VErrors(ve.Translate(v.t))}
	}
	return nil
}
