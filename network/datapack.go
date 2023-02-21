package network

import (
	"bytes"
	"encoding/binary"
	"errors"
	"gee/iface"
	"gee/utils"
)

type DataPack struct {
}

func NewDataPack() iface.IDataPack {
	return &DataPack{}
}

func (p *DataPack) GetHeaderLength() uint32 {
	return 8
}

// Pack 打包消息
// 包括消息长度、消息🆔、消息内容
func (p *DataPack) Pack(message iface.IMessage) ([]byte, error) {
	// 创建缓冲区
	buf := bytes.NewBuffer([]byte{})

	// 1.将数据长度写入缓冲
	if err := binary.Write(buf, binary.LittleEndian, message.GetLength()); err != nil {
		return nil, err
	}

	// 2.将消息🆔写入缓冲
	if err := binary.Write(buf, binary.LittleEndian, message.GetID()); err != nil {
		return nil, err
	}

	// 3.将数据写入缓冲
	if err := binary.Write(buf, binary.LittleEndian, message.GetData()); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Unpack 消息拆包
func (p *DataPack) Unpack(data []byte) (iface.IMessage, error) {
	reader := bytes.NewReader(data)
	message := &Message{}

	// 1.header处理
	{
		if utils.TCP.MaxPackageSize > 0 && int(message.Length) > utils.TCP.MaxPackageSize {
			return nil, errors.New("too large message received")
		}

		// 1.读消息长度
		if err := binary.Read(reader, binary.LittleEndian, &message.Length); err != nil {
			return nil, err
		}

		// 2.读消息🆔
		if err := binary.Read(reader, binary.LittleEndian, &message.ID); err != nil {
			return nil, err
		}
	}

	return message, nil
}
