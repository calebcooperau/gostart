package dtos

import (
	"github.com/google/uuid"
)

type UserDetailApiDto struct {
	ID           uuid.UUID
	Email        string
	FirstName    string
	LastName     string
	MobileNumber *string
}
