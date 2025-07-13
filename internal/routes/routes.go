package routes

import (
	"log"

	"gostart/internal/domain"
	authRoutes "gostart/internal/domain/authentication/api/routes"
	userRoutes "gostart/internal/domain/user/api/routes"
	"gostart/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(engine *gin.Engine, db *pgxpool.Pool, repos *domain.Repositories, middlewares *middleware.Middlewares, logger *log.Logger) {
	router := engine
	router.Use(middlewares.AuthenticationMiddleware.Authenticate())
	{
		authRoutes.RegisterAuthenticationRoutes(router, repos.AuthenticationRepository, logger)
		userRoutes.RegisterUserRoutes(router, db, repos.UserRepository, middlewares.AuthenticationMiddleware, logger)
	}
}
