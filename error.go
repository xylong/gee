package gee

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ErrorHandler 错误处理
func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error": err,
				})
			}
		}()
		context.Next()
	}
}

// Error 错误信息
func Error(err error, message ...string) {
	if err == nil {
		return
	}
	msg := err.Error()
	if len(message) > 0 {
		msg = message[0]
	}
	panic(msg)
}
