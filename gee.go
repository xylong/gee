package gee

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
)

// Gee
type Gee struct {
	*gin.Engine
	group *gin.RouterGroup
	dba   interface{}
}

// Init 初始化
func Init() *Gee {
	gee := &Gee{Engine: gin.New()}
	gee.Use(ErrorHandler())
	return gee
}

// Go 启动
func (gee *Gee) Go() {
	if err := gee.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// Mount 挂载控制器路由
// group 路由分组
// controllers 控制器
func (gee *Gee) Mount(group string, controllers ...Controller) *Gee {
	gee.group = gee.Group(group)
	for _, controller := range controllers {
		controller.Build(gee)
		value := reflect.ValueOf(controller).Elem()
		if value.NumField() > 0 {
			if gee.dba != nil {
				value.Field(0).Set(reflect.New(value.Field(0).Type().Elem()))
				value.Field(0).Elem().Set(reflect.ValueOf(gee.dba).Elem())
			}
		}
	}
	return gee
}

// Handle 重载gin的Handle方法
func (gee *Gee) Handle(httpMethod, relativePath string, handler interface{}) *Gee {
	if handle := Convert(handler); handle != nil {
		gee.group.Handle(httpMethod, relativePath, handle)
	}
	return gee
}

// Attach 附加中间件
func (gee *Gee) Attach(middleware Middleware) *Gee {
	gee.Use(func(context *gin.Context) {
		if err := middleware.OnRequest(context); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			context.Next()
		}
	})
	return gee
}

// DB 设置数据库🔗对象
func (gee *Gee) DB(dba interface{}) *Gee {
	gee.dba = dba
	return gee
}
