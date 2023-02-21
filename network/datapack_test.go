package network

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {
	// 模拟服务器
	listener, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		t.Logf("[server listen]%s", err.Error())
		return
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				t.Logf("[server accept]%s", err.Error())
				return
			}

			// 拆包
			go func(conn net.Conn) {
				dp := NewDataPack()

				for {
					// 1.先读出header中的数据长度
					header := make([]byte, dp.GetHeaderLength())
					if _, err := io.ReadFull(conn, header); err != nil {
						t.Logf("[conn read1]%s", err.Error())
						break
					}

					// 2.再根据读出的长度读取消息内容
					message, err := dp.Unpack(header)
					if err != nil {
						t.Logf("[unpack]%s", err.Error())
						return
					}
					if message.GetLength() > 0 {
						data := make([]byte, message.GetLength())
						if _, err := io.ReadFull(conn, data); err != nil {
							t.Logf("[conn read2]%s", err.Error())
							return
						}

						message.(*Message).Data = data
						fmt.Println("received", "id=", message.GetID(), " length=", message.GetLength(), " data=", string(message.GetData()))
					}
				}
			}(conn)
		}
	}()

	// 模拟客户端
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		t.Logf("[dail]%s", err.Error())
		return
	}

	// 模拟粘包
	dp := NewDataPack()
	// 将两个包在一起发
	{
		msg1 := &Message{1, 5, []byte("hello")}
		msg2 := &Message{2, 3, []byte("gee")}

		data, err := dp.Pack(msg1)
		if err != nil {
			t.Logf("[first pack]%s", err.Error())
			return
		}
		data2, err := dp.Pack(msg2)
		if err != nil {
			t.Logf("[second  pack]%s", err.Error())
			return
		}

		// 将包粘在一起
		data = append(data, data2...)

		// 一次发送给服务器
		if _, err = conn.Write(data); err != nil {
			t.Log(err)
			return
		}
	}

	select {}
}
