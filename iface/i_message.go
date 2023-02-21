package iface

// IMessage 请求消息
type IMessage interface {
	// GetID 获取🆔
	GetID() uint32

	// SetID 设置🆔
	SetID(uint32)

	// GetLength 获取消息长度
	GetLength() uint32

	// SetLength 设置消息长度
	SetLength(uint32)

	// GetData 获取消息
	GetData() []byte

	// SetData 设置消息
	SetData([]byte)
}
