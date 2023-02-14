package main

import (
	"fmt"
	"gee/iface"
	"gee/network"
)

type pingRouter struct {
	network.BaseRouter
}

func (r *pingRouter) Before(request iface.IRequest) {
	fmt.Println("before")
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping\n")); err != nil {
		fmt.Printf("before ping error:%s", err.Error())
	}

}
func (r *pingRouter) Handle(request iface.IRequest) {
	fmt.Println("handling")
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("ping\n")); err != nil {
		fmt.Printf("handling ping error:%s", err.Error())
	}

}
func (r *pingRouter) After(request iface.IRequest) {
	fmt.Println("after")
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping\n")); err != nil {
		fmt.Printf("after ping error:%s", err.Error())
	}

}

func main() {
	s := network.NewServer("Gee")
	s.Route(&pingRouter{})
	s.Run()
}
