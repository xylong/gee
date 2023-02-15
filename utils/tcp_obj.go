package utils

import (
	"gee/iface"
)

// TCPObject tcp连接操作对象
type TCPObject struct {
	// server
	Server iface.IServer
	Host   string
	Port   int
	Name   string

	Version        string // 版本号
	MaxConn        int    // 最大连接数
	MaxPackageSize int    // 数据包最大值
}

// Reload 重载配置
func (o *TCPObject) Reload() {
	/*
		data, err := os.ReadFile("config.json")
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(data, o); err != nil {
			panic(err)
		}
	*/
}

// TCP tcp对象
var TCP *TCPObject

func init() {
	TCP = &TCPObject{
		Host:           "0.0.0.0",
		Port:           8080,
		Name:           "Gee",
		Version:        "v0.0.1",
		MaxConn:        100,
		MaxPackageSize: 1024,
	}

	TCP.Reload()
}
