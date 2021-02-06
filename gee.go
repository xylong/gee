package gee

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Gee
type Gee struct {
	*gin.Engine
	group       *gin.RouterGroup
	beanFactory *BeanFactory
}

// Init 初始化
func Init() *Gee {
	gee := &Gee{
		Engine:      gin.New(),
		beanFactory: NewBeanFactory(),
	}
	gee.Use(ErrorHandler())
	// 加载配置
	config := InitConfig()
	gee.beanFactory.setBean(config)
	// 加载模板
	if config.Server.Html != "" {
		gee.LoadHTMLGlob(config.Server.Html)
	}

	return gee
}

// Go 启动
func (gee *Gee) Go() {
	if err := gee.Run(fmt.Sprintf(":%d", InitConfig().Server.Port)); err != nil {
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
		gee.beanFactory.inject(controller)
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

func (gee *Gee) Beans(beans ...interface{}) *Gee {
	gee.beanFactory.setBean(beans)
	return gee
}
