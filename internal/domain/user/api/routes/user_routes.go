package routes

import (
	"log"

	"gostart/internal/domain/user/api/delete"
	"gostart/internal/domain/user/api/detail"
	"gostart/internal/domain/user/api/update"
	"gostart/internal/domain/user/data"
	"gostart/internal/domain/user/repository"
	"gostart/internal/middleware/authentication"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterUserRoutes(router *gin.Engine, db *pgxpool.Pool, repo repository.UserRepository, authMiddleware authentication.AuthenticationMiddleware, logger *log.Logger) {
	queries := data.New(db)
	// Set up handlers
	detailHandler := detail.NewUserDetailHandler(queries, logger)
	updateHandler := update.NewUserUpdateHandler(repo, logger)
	deleteHandler := delete.NewUserDeleteHandler(repo, logger)

	// Set up routes
	userRoutes := router.Group("/user")
	userRoutes.Use(authMiddleware.RequireAuthUser())
	{
		userRoutes.GET("/:id", detailHandler.GetUserByID)
		userRoutes.PUT("/:id", updateHandler.UpdateUser)
		userRoutes.DELETE("/:id", deleteHandler.DeleteUser)
	}
}
