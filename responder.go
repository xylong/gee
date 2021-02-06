package gee

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

var Responders []Responder

func init() {
	Responders = []Responder{
		new(StringResponder),
		new(ObjectResponder),
		new(SliceResponder),
		new(ViewResponder),
	}
}

// Responder 响应者
type Responder interface {
	Respond() gin.HandlerFunc
}

// Convert 转化
// 将自定义Handler转换成gin的HandlerFunc
func Convert(handler interface{}) gin.HandlerFunc {
	value := reflect.ValueOf(handler)

	for _, responder := range Responders {
		ptr := reflect.ValueOf(responder).Elem()
		if value.Type().ConvertibleTo(ptr.Type()) {
			ptr.Set(value)
			return ptr.Interface().(Responder).Respond()
		}
	}

	return nil
}

// StringResponder 字符串响应
type StringResponder func(*gin.Context) string

// Respond 响应字符串
func (r StringResponder) Respond() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusOK, r(context))
	}
}

// ObjectResponder 实体响应
type ObjectResponder func(*gin.Context) Object

// Respond 响应实体
func (r ObjectResponder) Respond() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, r(context))
	}
}

// SliceResponder 切片响应
type SliceResponder func(*gin.Context) Objects

// Respond 切片实体
func (r SliceResponder) Respond() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Content-type", "application/json")
		_, _ = context.Writer.WriteString(string(r(context)))
	}
}

type (
	// View 视图
	View          string
	ViewResponder func(ctx *gin.Context) View
)

// Respond 视图响应
func (r ViewResponder) Respond() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.HTML(http.StatusOK, fmt.Sprintf("%s.html", string(r(context))), nil)
	}
}
