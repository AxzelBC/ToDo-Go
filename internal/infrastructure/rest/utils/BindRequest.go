package utils

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
)

func BindToJson(ctx *gin.Context, request any) (err error) {
	buffer := make([]byte, 5120)
	num, _ := ctx.Request.Body.Read(buffer)
	requestBody := string(buffer[0:num])
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(requestBody)))
	err = ctx.ShouldBindJSON(request)
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(requestBody)))
	return
}
