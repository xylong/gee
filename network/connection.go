package network

import (
	"fmt"
	"gee/iface"
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

	// 业务处理方法
	handle iface.HandleFunc

	// 连接退出通知
	ExitChan chan struct{}
}

// NewConnection 新建连接
func NewConnection(conn *net.TCPConn, id uint32, callback iface.HandleFunc) iface.IConnection {
	return &Connection{
		Conn:     conn,
		ID:       id,
		handle:   callback,
		isClosed: false,
		ExitChan: make(chan struct{}, 1),
	}
}

func (c *Connection) read() {
	defer c.Stop()

	for {
		buf := make([]byte, 1024)
		length, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("receive error:", err.Error())
			continue
		}

		if err := c.handle(c.Conn, buf, length); err != nil {
			fmt.Printf("ConnID:%d handle error:%s", c.ID, err.Error())
			break
		}
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
