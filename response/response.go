package response

import (
	"github.com/gin-gonic/gin"
)

const (
	successCode = 0
	errorCode   = 1
)

type response struct {
	status_code int
	status_msg  string
}

func Response(ctx *gin.Context, httpStatus int, v interface{}) {
	ctx.JSON(httpStatus, v)
}

func Success(ctx *gin.Context, msg string, v interface{}) {
	if v == nil {
		Response(ctx, 200, response{successCode, msg})
		return
	}
}
func Fail(ctx *gin.Context, msg string, v interface{}) {
	if v == nil {
		Response(ctx, 400, response{errorCode, msg})
		return
	}
	//
}
