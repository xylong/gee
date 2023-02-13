package iface

// IRequest 请求接口
type IRequest interface {
	// GetConnection 获取当前连接
	GetConnection() IConnection

	// Data 获取数据
	Data() []byte
}
