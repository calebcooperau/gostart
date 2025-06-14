package authentication

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func Provider(ctx *gin.Context) {
	// try to get the user without re-authenticating
	provider := ctx.Param("provider")

	newCtx := context.WithValue(ctx.Request.Context(), "provider", provider)
	ctx.Request = ctx.Request.WithContext(newCtx)
	user, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
	} else {
		fmt.Print(user)
		http.Redirect(ctx.Writer, ctx.Request, "http://localhost:4200", http.StatusFound)
	}
}
