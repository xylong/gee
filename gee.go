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
	props []interface{}
}

// Init 初始化
func Init() *Gee {
	gee := &Gee{
		Engine: gin.New(),
		props:  make([]interface{}, 0),
	}
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
		gee.setProp(controller)
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
	gee.props = append(gee.props, dba)
	return gee
}

// setProp 设置控制器属性
func (gee *Gee) setProp(controller Controller) {
	ptr := reflect.ValueOf(controller).Elem()
	for i := 0; i < ptr.NumField(); i++ {
		field := ptr.Field(i)
		// 判断控制器属性是否已经实例化
		if !field.IsNil() || field.Kind() != reflect.Ptr {
			continue
		}
		// 创建属性
		if p := gee.getProp(field.Type()); p != nil {
			field.Set(reflect.New(field.Type().Elem()))
			field.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}

// getProp 获取控制器属性
func (gee *Gee) getProp(p reflect.Type) interface{} {
	for _, prop := range gee.props {
		if p == reflect.TypeOf(prop) {
			return prop
		}
	}
	return nil
}
