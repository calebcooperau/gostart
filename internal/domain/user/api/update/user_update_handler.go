package update

import (
	"encoding/json"
	"log"
	"net/http"

	"gostart/internal/domain/user/api/update/dtos"
	"gostart/internal/domain/user/repository"
	"gostart/internal/utilities"

	"github.com/gin-gonic/gin"
)

type UserUpdateHandler struct {
	repository repository.UserRepository
	logger     *log.Logger
}

func NewUserUpdateHandler(repository repository.UserRepository, logger *log.Logger) *UserUpdateHandler {
	return &UserUpdateHandler{
		repository: repository,
		logger:     logger,
	}
}

func (handler UserUpdateHandler) UpdateUser(ctx *gin.Context) {
	userID, err := utilities.ReadIDParam(ctx)
	if err != nil {
		handler.logger.Printf("ERROR: readIDParam: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Sent"})
		return
	}

	// build command
	var userUpdateApiDto dtos.UserUpdateApiDto
	err = json.NewDecoder(ctx.Request.Body).Decode(&userUpdateApiDto)
	if err != nil {
		handler.logger.Printf("Error: decodeUserUpdateApiDto: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Sent"})
		return
	}

	// validate
	err = userUpdateApiDto.ValidateApiDto()
	if err != nil {
		handler.logger.Printf("ERROR: validateUserUpdateApiDto: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Sent"})
		return
	}

	command := UserUpdateCommand{
		ID:           userID,
		Email:        userUpdateApiDto.Email,
		FirstName:    userUpdateApiDto.FirstName,
		LastName:     userUpdateApiDto.LastName,
		MobileNumber: userUpdateApiDto.MobileNumber,
	}

	user, err := handler.repository.FindUserByID(ctx.Request.Context(), command.ID)
	if err != nil {
		handler.logger.Printf("Error: repositoryGetUserByID: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	user, err = user.Update(command.Email, command.FirstName, command.LastName, command.MobileNumber)
	if err != nil {
		handler.logger.Printf("Error: modelUserUpdate: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	user, err = handler.repository.UpdateUser(ctx.Request.Context(), user)
	if err != nil {
		handler.logger.Printf("Error: repositoryUpdateUser: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"User": user})
}
