package health_check

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetApplicationHealthCheck(ctx *gin.Context) {
	fmt.Fprintf(ctx.Writer, "Status is available\n")
}
