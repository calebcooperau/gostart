package routes

import (
	"log"

	"gostart/internal/domain/user"
	"gostart/internal/features/user/delete"
	"gostart/internal/features/user/detail"
	"gostart/internal/features/user/update"
	"gostart/internal/middleware/authentication"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, repo user.UserRepository, authMiddleware authentication.AuthenticationMiddleware, logger *log.Logger) {
	// Set up handlers
	detailHandler := detail.NewUserDetailHandler(repo, logger)
	updateHandler := update.NewUserUpdateHandler(repo, logger)
	deleteHandler := delete.NewUserDeleteHandler(repo, logger)

	// Set up routes
	userRoutes := router.Group("/user")
	userRoutes.Use(authMiddleware.RequireAuthUser())
	{
		userRoutes.GET("/:id", detailHandler.GetUserById)
		userRoutes.PUT("/:id", updateHandler.UpdateUser)
		userRoutes.DELETE("/:id", deleteHandler.DeleteUser)
	}
}
