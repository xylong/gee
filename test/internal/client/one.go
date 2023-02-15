package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Printf("dail error:%s\n", err.Error())
	}

	for {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			fmt.Printf("write error:%s\n", err.Error())
			return
		}

		buf := make([]byte, 1024)
		if length, err := conn.Read(buf); err != nil {
			fmt.Printf("read error::%s", err.Error())
		} else {
			fmt.Printf("received[%d]:%s\n", length, string(buf))
		}

		time.Sleep(time.Second * 2)
	}
}
