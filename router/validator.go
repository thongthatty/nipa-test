package router

import "gopkg.in/go-playground/validator.v9"

// NewValidator create struct of Validator
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validator struct
type Validator struct {
	validator *validator.Validate
}

// Validate function that validate interface of input
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
