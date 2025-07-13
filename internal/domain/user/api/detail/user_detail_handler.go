package detail

import (
	"log"
	"net/http"

	"gostart/internal/domain/user/api/detail/dtos"
	"gostart/internal/domain/user/data"
	"gostart/internal/utilities"

	"github.com/gin-gonic/gin"
)

type UserDetailHandler struct {
	queries *data.Queries
	logger  *log.Logger
}

func NewUserDetailHandler(queries *data.Queries, logger *log.Logger) *UserDetailHandler {
	return &UserDetailHandler{
		queries: queries,
		logger:  logger,
	}
}

func (handler UserDetailHandler) GetUserByID(ctx *gin.Context) {
	userID, err := utilities.ReadIDParam(ctx)
	if err != nil {
		handler.logger.Printf("ERROR: ReadIDParam: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	query := UserDetailQuery{
		ID: userID,
	}

	user, err := handler.queries.GetUserDetailByID(ctx.Request.Context(), query.ID)
	if err != nil {
		handler.logger.Printf("ERROR: repositoryGetUserByID: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	userApiDto := dtos.UserDetailApiDto{
		ID:           user.ID,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		MobileNumber: user.MobileNumber,
	}

	ctx.JSON(http.StatusOK, gin.H{"User": userApiDto})
}
