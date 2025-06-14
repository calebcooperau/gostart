package authentication

import (
	"time"

	"github.com/google/uuid"
)

type AuthUser struct {
	Id        uuid.UUID
	Email     string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
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
