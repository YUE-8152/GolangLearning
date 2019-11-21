package commutils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  http.StatusText(200),
		"data": data,
	})
}

func FailedResponse(ctx *gin.Context, data interface{}, msg string) {
	statusCode := ctx.Writer.Status()
	if msg == "" {
		msg = http.StatusText(statusCode)
	}
	ctx.JSON(statusCode, gin.H{
		"code": statusCode,
		"msg":  msg,
		"data": data,
	})
}
