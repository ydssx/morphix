package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	ERROR      = -1
	SUCCESS    = 0
	ErrorMsg   = "操作失败"
	SuccessMsg = "操作成功"
)

func result(c *gin.Context, httpCode int, code int, data interface{}, msg string) {
	c.JSON(httpCode, response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func OK(c *gin.Context) {
	result(c, http.StatusOK, SUCCESS, nil, SuccessMsg)
}

func OKWithData(c *gin.Context, data interface{}) {
	result(c, http.StatusOK, SUCCESS, data, SuccessMsg)
}

func FailWithMsg(c *gin.Context, msg string) {
	result(c, http.StatusBadRequest, ERROR, nil, msg)
}

func FailWithError(c *gin.Context, err error) {
	result(c, http.StatusBadRequest, ERROR, nil, wrapValidateErrMsg(err))
}

func wrapValidateErrMsg(err error) (msg string) {
	switch v := err.(type) {
	case *json.UnmarshalTypeError:
		msg = fmt.Sprintf("请求参数`%s`类型错误，应为%s类型", v.Field, v.Type.Name())
	case validator.ValidationErrors:
		for _, e := range v {
			msg += fmt.Sprintf("缺少必要参数：`%s`", strings.ToLower(e.Field()))
		}
	default:
		msg = err.Error()
	}
	return
}
