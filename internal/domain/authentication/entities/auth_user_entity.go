package entities

import "github.com/google/uuid"

type AuthUser struct {
	ID           uuid.UUID
	Email        string
	FirstName    string
	LastName     string
	MobileNumber *string
}

var AnonymousUser = &AuthUser{}

func (usr *AuthUser) IsAnonymous() bool {
	return usr == AnonymousUser
}

func Create(email string, firstName string, lastName string) (*AuthUser, error) {
	user := AuthUser{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
	return &user, nil
}
