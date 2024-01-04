package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ps struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	OriginUrl string      `json:"originUrl"`
}

func Succ(ctx *gin.Context, data interface{}) {
	resp := ps{
		Code:      0,
		Message:   "OK",
		Data:      data,
		OriginUrl: ctx.Request.URL.Path,
	}
	ctx.Set("succ_response", resp)
	ctx.JSON(http.StatusOK, resp)
}

func Err(ctx *gin.Context, ErrCode int, messge string) {
	resp := ps{
		Code:      ErrCode,
		Error:     "error some",
		Message:   messge,
		OriginUrl: ctx.Request.URL.Path,
	}
	ctx.Set("err_response", resp)
	ctx.JSON(http.StatusOK, resp)
}
