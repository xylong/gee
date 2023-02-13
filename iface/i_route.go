package iface

type IRouter interface {
	// Before 业务执行前的钩子方法
	Before(IRequest)

	// Handle 业务处理钩子方法
	Handle(IRequest)

	// After 业务执行后的钩子方法
	After(IRequest)
}
