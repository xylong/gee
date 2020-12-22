package gee

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Gee
type Gee struct {
	*gin.Engine
	group *gin.RouterGroup
}

// Init 初始化
func Init() *Gee {
	return &Gee{Engine: gin.New()}
}

// Go 启动
func (g *Gee) Go() {
	if err := g.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// Mount 挂载控制器路由
// group 路由分组
// controllers 控制器
func (g *Gee) Mount(group string, controllers ...Controller) *Gee {
	g.group = g.Group(group)
	for _, controller := range controllers {
		controller.Build(g)
	}
	return g
}

// Handle 重载gin的Handle方法
func (g *Gee) Handle(httpMethod, relativePath string, handler interface{}) *Gee {
	if handle := Convert(handler); handle != nil {
		g.group.Handle(httpMethod, relativePath, handle)
	}
	return g
}

// Attach 附加中间件
func (g *Gee) Attach(middleware Middleware) *Gee {
	g.Use(func(context *gin.Context) {
		if err := middleware.OnRequest(context); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			context.Next()
		}
	})
	return g
}
