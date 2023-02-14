package network

import (
	"fmt"
	"gee/iface"
	"net"
)

// Server 服务
type Server struct {
	// 服务名称
	Name string

	// ip版本
	IPVersion string

	// ip地址
	IP string

	// 端口
	Port int

	// 路由
	Router iface.IRouter
}

// NewServer 新建服务
func NewServer(name string) iface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8080,
	}
}

func (s *Server) Start() {
	go func() {
		// 1.获取tcp地址
		tcpAddr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve error: ", err.Error())
			return
		}

		// 2.监听tcp地址
		listener, err := net.ListenTCP(s.IPVersion, tcpAddr)
		if err != nil {
			fmt.Println("listen error: ", err.Error())
			return
		}
		fmt.Printf("listening at [%s:%d]", s.IP, s.Port)

		// 3.⌛️客户端连接，处理业务消息
		var connID uint32

		for {
			tcpConn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept error: ", err.Error())
				continue
			}

			conn := NewConnection(tcpConn, connID, s.Router)
			conn.Start()

			connID++
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Run() {
	s.Start()

	// ⌛️
	select {}
}

// Route 设置路由
func (s *Server) Route(router iface.IRouter) {
	s.Router = router
}
