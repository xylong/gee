package iface

import "net"

// IConnection 连接模块抽象接口
type IConnection interface {
	// Start 启动连接
	Start()

	// Stop 停止连接
	Stop()

	// GetTCPConnection 获取socket
	GetTCPConnection() *net.TCPConn

	// GetConnID 获取连接🆔
	GetConnID() uint32

	// GetRemoteAddr 获取远程客户段tcp状态
	// ip、port
	GetRemoteAddr() *net.TCPAddr

	// Send 发送数据
	Send([]byte) error
}

// HandleFunc 业务处理函数
type HandleFunc func(*net.TCPConn, []byte, int) error
