package routes

import (
	"log"

	"gostart/internal/domain/authentication"

	"gostart/internal/features/authentication/logout"
	"gostart/internal/features/authentication/provider"
	"gostart/internal/features/authentication/signin"

	"github.com/gin-gonic/gin"
)

func RegisterAuthenticationRoutes(router *gin.Engine, authenticationRepo authentication.AuthenticationRepository, logger *log.Logger) {
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
