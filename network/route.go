package network

import "gee/iface"

// BaseRouter 基础路由
// 方便有的router不希望实现router接口所有方法，可用此基础路由嵌套，然后重写需要的方法即可
type BaseRouter struct {
}

func (r *BaseRouter) Before(request iface.IRequest) {}
func (r *BaseRouter) Handle(request iface.IRequest) {}
func (r *BaseRouter) After(request iface.IRequest)  {}
