package iface

// IDataPack 封包、拆包
type IDataPack interface {
	// GetHeaderLength 获取包头长度
	GetHeaderLength() uint32

	// Pack 打包消息
	Pack(IMessage) ([]byte, error)

	// Unpack 消息拆包
	Unpack([]byte) (IMessage, error)
}
