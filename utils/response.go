package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status     string      `json:"status"`
	BaseStatus string      `json:"baseStatus"`
	Message    string      `json:"msg"`
	Data       interface{} `json:"data,omitempty"`
}

func BadRequestResp(ctx *gin.Context, message string) {
	ctx.Error(errors.New(message))

	ctx.JSON(http.StatusBadRequest, &Response{
		Status:     "400",
		BaseStatus: "BAD_REQUEST",
		Message:    message,
		Data:       nil,
	})
}

func OkResp(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, &Response{
		Status:     "0",
		BaseStatus: "OK",
		Message:    message,
		Data:       data,
	})
}
