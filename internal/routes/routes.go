package routes

import (
	"log"

	"gostart/internal/domain"
	authRoutes "gostart/internal/features/authentication/routes"
	health_check_handler "gostart/internal/features/health_check"
	userRoutes "gostart/internal/features/user/routes"
	"gostart/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine, repos *domain.Repositories, middlewares *middleware.Middlewares, logger *log.Logger) {
	router := engine
	router.GET("/health", health_check_handler.GetApplicationHealthCheck)
	router.Use(middlewares.AuthenticationMiddleware.Authenticate())
	{
		authRoutes.RegisterAuthenticationRoutes(router, repos.AuthenticationRepository, logger)
		userRoutes.RegisterUserRoutes(router, repos.UserRepository, middlewares.AuthenticationMiddleware, logger)
	}
}
