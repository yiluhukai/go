package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code": 0, //0表示成功,其他表示失败
	"message":"success"， //用来描述失败的原因
	"data":{
	}
}
*/
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseError(context *gin.Context, code int) {
	responseData := ResponseData{
		Code:    code,
		Message: GetMessage(code),
	}
	context.JSON(http.StatusOK, responseData)
}

func ResponseSuccess(context *gin.Context, data interface{}) {
	responseData := ResponseData{
		Code:    ErrCodeSuccess,
		Data:    data,
		Message: "请求成功",
	}
	context.JSON(http.StatusOK, responseData)
}
