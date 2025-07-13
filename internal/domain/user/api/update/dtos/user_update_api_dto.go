package dtos

import "github.com/go-playground/validator"

type UserUpdateApiDto struct {
	Email        string `json:"email" validate:"required"`
	FirstName    string `json:"firstName" validate:"required"`
	LastName     string `json:"lastName" validate:"required"`
	MobileNumber string `json:"mobileNumber" `
}

func (dto *UserUpdateApiDto) ValidateApiDto() error {
	validate := validator.New()
	return validate.Struct(dto)
}
