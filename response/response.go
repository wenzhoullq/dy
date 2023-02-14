package response

import (
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
)

const (
	successCode = 0
	errorCode   = -1
)

type Resp struct {
	Statuscode int    `json:"status_code"`
	Statusmsg  string `json:"status_msg"`
}

func checkResponese(v interface{}) {
	getValue := reflect.ValueOf(v)
	Msg := getValue.Elem().FieldByName("StatusMsg")
	if !Msg.CanSet() {
		log.Println("cant set msg")
	}
	Code := getValue.Elem().FieldByName("StatusCode")
	if !Code.CanSet() {
		log.Println("cant set StatusCode")
	}
}

func response(ctx *gin.Context, httpStatus int, v interface{}) {
	ctx.JSON(httpStatus, v)
}

func Success(ctx *gin.Context, msg string, v interface{}) {
	if v == nil {
		response(ctx, 200, Resp{successCode, msg})
		return
	}
	response(ctx, 200, v)
}
func Fail(ctx *gin.Context, msg string, v interface{}) {
	if v == nil {
		response(ctx, 200, Resp{errorCode, msg})
		return
	}
	//checkResponese(v)
	response(ctx, 200, v)
}
