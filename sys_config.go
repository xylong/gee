package gee

import (
	"github.com/xylong/gee/pkg/util"
	"gopkg.in/yaml.v2"
	"log"
)

type ServerConfig struct {
	Port int32
	Name string
}

type SysConfig struct {
	Server *ServerConfig
	Config UserConfig
}

type UserConfig map[string]interface{}

func NewSysConfig() *SysConfig {
	return &SysConfig{
		Server: &ServerConfig{
			Port: 8080,
			Name: "gee",
		}}
}

func InitConfig() *SysConfig {
	config := NewSysConfig()
	if b := util.LoadConfigFile(); b != nil {
		if err := yaml.Unmarshal(b, config); err != nil {
			log.Fatal(err)
		}
	}
	return config
}
