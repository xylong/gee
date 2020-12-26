package util

import (
	"io/ioutil"
	"log"
	"os"
)

// LoadConfigFile 加载配置文件
func LoadConfigFile() []byte {
	dir, _ := os.Getwd()
	file := dir + "/app.yaml"
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}
	return b
}
