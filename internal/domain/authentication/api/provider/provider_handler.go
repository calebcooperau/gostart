package provider

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type ProviderHandler struct {
	logger *log.Logger
}

func NewProviderHandler(logger *log.Logger) *ProviderHandler {
	return &ProviderHandler{
		logger: logger,
	}
}

func (handler ProviderHandler) GetProvider(ctx *gin.Context) {
	provider := ctx.Param("provider")

	newCtx := context.WithValue(ctx.Request.Context(), "provider", provider)
	ctx.Request = ctx.Request.WithContext(newCtx)
	_, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
	} else {
		ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:4200")
	}
}
