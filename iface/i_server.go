package iface

// IServer 服务器接口
type IServer interface {
	// Start 启动服务
	Start()

	// Stop 停止服务
	Stop()

	// Run 运行服务
	Run()

	// Route 设置路由
	Route(IRouter)
}
