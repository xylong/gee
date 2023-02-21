package network

type Message struct {
	// 消息🆔
	ID uint32

	// 消息长度
	Length uint32

	// 消息内容
	Data []byte
}

func (m *Message) GetID() uint32 {
	return m.ID
}

func (m *Message) SetID(id uint32) {
	m.ID = id
}

func (m *Message) GetLength() uint32 {
	return m.Length
}

func (m *Message) SetLength(length uint32) {
	m.Length = length
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}
