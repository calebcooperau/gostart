package middleware

import (
	"gostart/internal/domain"
	"gostart/internal/middleware/authentication"
)

type Middlewares struct {
	AuthenticationMiddleware authentication.AuthenticationMiddleware
}

func RegisterMiddlewares(repositories *domain.Repositories) *Middlewares {
	authenticationMiddleware := authentication.AuthenticationMiddleware{AuthenticationRepository: repositories.AuthenticationRepository}

	middlewares := &Middlewares{
		AuthenticationMiddleware: authenticationMiddleware,
	}

	return middlewares
}
