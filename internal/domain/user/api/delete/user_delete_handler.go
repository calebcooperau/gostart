package delete

import (
	"database/sql"
	"log"
	"net/http"

	"gostart/internal/domain/user/repository"
	"gostart/internal/utilities"

	"github.com/gin-gonic/gin"
)

type UserDeleteHandler struct {
	repository repository.UserRepository
	logger     *log.Logger
}

func NewUserDeleteHandler(repository repository.UserRepository, logger *log.Logger) *UserDeleteHandler {
	return &UserDeleteHandler{
		repository: repository,
		logger:     logger,
	}
}

func (handler UserDeleteHandler) DeleteUser(ctx *gin.Context) {
	userID, err := utilities.ReadIDParam(ctx)
	if err != nil {
		handler.logger.Printf("ERROR: readIDParam: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	UserDeleteCommand := UserDeleteCommand{
		ID: userID,
	}
	user, err := handler.repository.FindUserByID(ctx.Request.Context(), UserDeleteCommand.ID)
	if err != nil {
		handler.logger.Printf("ERROR: repositoryGetUserByID: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	err = user.CanDelete()
	if err != nil {
		handler.logger.Printf("ERROR: userCanDelete: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	err = handler.repository.DeleteUser(ctx.Request.Context(), user)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	if err != nil {
		handler.logger.Printf("ERROR repositoryDeleteUser: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
}
