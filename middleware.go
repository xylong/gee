package gee

import "github.com/gin-gonic/gin"

// Middleware 中间件
type Middleware interface {
	OnRequest(*gin.Context) error
}
