package detail

import (
	"log"
	"net/http"

	"gostart/internal/domain/user"
	"gostart/internal/features/user/detail/dtos"
	"gostart/internal/utilities"

	"github.com/gin-gonic/gin"
)

type UserDetailHandler struct {
	repository user.UserRepository
	logger     *log.Logger
}

func NewUserDetailHandler(repository user.UserRepository, logger *log.Logger) *UserDetailHandler {
	return &UserDetailHandler{
		repository: repository,
		logger:     logger,
	}
}

func (handler UserDetailHandler) GetUserById(ctx *gin.Context) {
	userId, err := utilities.ReadIDParam(ctx)
	if err != nil {
		handler.logger.Printf("ERROR: ReadIDParam: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User Id"})
		return
	}

	query := UserDetailQuery{
		Id: userId,
	}

	user, err := handler.repository.GetUserById(query.Id)
	if err != nil {
		handler.logger.Printf("ERROR: repositoryGetUserById: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	userApiDto := dtos.UserDetailApiDto{
		Id:        user.Id,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"User": userApiDto})
}
