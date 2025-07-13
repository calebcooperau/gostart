package domain

import (
	"github.com/jackc/pgx/v5/pgxpool"
	authentication "gostart/internal/domain/authentication/repository"
	user "gostart/internal/domain/user/repository"
)

type Repositories struct {
	UserRepository           user.UserRepository
	AuthenticationRepository authentication.AuthenticationRepository
}

func RegisterRepositories(db *pgxpool.Pool) *Repositories {
	userRepository := user.NewUserSqlRepository(db)
	authenticationRepository := authentication.NewAuthenticationSqlRepository(db)
	return &Repositories{
		UserRepository:           userRepository,
		AuthenticationRepository: authenticationRepository,
	}
}
