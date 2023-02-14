package network

import "gee/iface"

// Request 请求
type Request struct {
	conn iface.IConnection
	data []byte
}

func (r *Request) GetConnection() iface.IConnection {
	return r.conn
}

func (r *Request) Data() []byte {
	return r.data
}
