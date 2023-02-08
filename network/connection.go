package network

import (
	"gee/iface"
	"net"
)

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

func NewConnection(conn *net.TCPConn, id uint32, callback iface.HandleFunc) iface.IConnection {
	return &Connection{
		Conn:     conn,
		ID:       id,
		handle:   callback,
		isClosed: false,
		ExitChan: make(chan struct{}, 1),
	}
}

func (c *Connection) Start() {

}

func (c *Connection) Stop() {

}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return nil
}

func (c *Connection) GetConnID() uint32 {
	return 0
}

func (c *Connection) GetRemoteAddr() *net.TCPAddr {
	return nil
}

func (c *Connection) Send(bytes []byte) error {
	return nil
}
