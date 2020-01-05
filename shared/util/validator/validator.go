package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator structure
type Validator struct {
	validator *validator.Validate
}

// New returns new validator
func New() *Validator {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return &Validator{
		validator: validate,
	}
}

// Validate validates request body
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
