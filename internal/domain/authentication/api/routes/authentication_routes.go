package routes

import (
	"log"

	"gostart/internal/domain/authentication/repository"

	"gostart/internal/domain/authentication/api/logout"
	"gostart/internal/domain/authentication/api/provider"
	"gostart/internal/domain/authentication/api/signin"

	"github.com/gin-gonic/gin"
)

func RegisterAuthenticationRoutes(router *gin.Engine, authenticationRepo repository.AuthenticationRepository, logger *log.Logger) {
	// Set up handlers
	signInHandler := signin.NewSignInHandler(authenticationRepo, logger)
	logoutHandler := logout.NewLogoutHandler(logger)
	providerHandler := provider.NewProviderHandler(logger)

	// Set up routes
	authRoutes := router.Group("/auth")
	authRoutes.GET("/:provider/callback", signInHandler.SignInCallback)
	authRoutes.GET("/logout/:provider", logoutHandler.Logout)
	authRoutes.GET("/:provider", providerHandler.GetProvider)
}
