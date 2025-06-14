package update

import "github.com/google/uuid"

type UserUpdateCommand struct {
	Id           uuid.UUID
	Email        string
	FirstName    string
	LastName     string
	MobileNumber string
}
