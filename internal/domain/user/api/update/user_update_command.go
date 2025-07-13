package update

import "github.com/google/uuid"

type UserUpdateCommand struct {
	ID           uuid.UUID
	Email        string
	FirstName    string
	LastName     string
	MobileNumber string
}
