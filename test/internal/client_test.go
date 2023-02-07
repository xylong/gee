package internal

import (
	"net"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatalf("client start error:%s\n", err.Error())
	}

	for {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			t.Logf("write error:%s\n", err.Error())
			return
		}

		buf := make([]byte, 1024)
		if length, err := conn.Read(buf); err != nil {
			t.Fatalf("read error::%s", err.Error())
		} else {
			t.Logf("received[%d]:%s\n", length, string(buf))
		}

		time.Sleep(time.Second * 2)
	}
}
