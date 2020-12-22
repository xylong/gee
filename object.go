package gee

import (
	"encoding/json"
	"log"
)

// Object 实体、对象
type Object interface {
	String() string
}

type Objects string

func MakeObjects(any interface{}) Objects {
	b, err := json.Marshal(any)
	if err != nil {
		log.Println(err)
	}
	return Objects(b)
}
