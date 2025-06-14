package domain

import (
	"database/sql"

	"gostart/internal/domain/authentication"
	"gostart/internal/domain/user"
)

type Repositories struct {
	UserRepository           user.UserRepository
	AuthenticationRepository authentication.AuthenticationRepository
}

func RegisterRepositories(db *sql.DB) *Repositories {
	userRepository := user.NewUserSqlRepository(db)
	authenticationRepository := authentication.NewAuthenticationSqlRepository(db)
	return &Repositories{
		UserRepository:           userRepository,
		AuthenticationRepository: authenticationRepository,
	}
}
