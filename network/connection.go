package network

import (
	"fmt"
	"gee/iface"
	"gee/utils"
	"net"
)

// Connection 连接模块
type Connection struct {
	// socket
	Conn *net.TCPConn

	// 连接🆔
	ID uint32

	// 是否关闭
	isClosed bool

	// 路由
	Router iface.IRouter

	// 连接退出通知
	ExitChan chan struct{}
}

// NewConnection 新建连接
func NewConnection(conn *net.TCPConn, id uint32, router iface.IRouter) iface.IConnection {
	return &Connection{
		Conn:     conn,
		ID:       id,
		isClosed: false,
		Router:   router,
		ExitChan: make(chan struct{}, 1),
	}
}

func (c *Connection) read() {
	defer c.Stop()

	for {
		buf := make([]byte, utils.TCP.MaxPackageSize)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("receive error:", err.Error())
			continue
		}

		req := &Request{
			conn: c,
			data: buf,
		}

		go func(request iface.IRequest) {
			c.Router.Before(request)
			c.Router.Handle(request)
			c.Router.After(request)
		}(req)
	}
}

func (c *Connection) Start() {
	go c.read()
}

func (c *Connection) Stop() {
	// 连接已关闭
	if c.isClosed {
		return
	}

	c.isClosed = true
	c.Conn.Close()
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(bytes []byte) error {
	return nil
}
