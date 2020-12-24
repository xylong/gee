package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"log"
)

type Xorm struct {
	*xorm.Engine
}

func NewXorm() *Xorm {
	engine, err := xorm.NewEngine("mysql", "root:root@/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	engine.DB().SetMaxOpenConns(5)
	engine.DB().SetMaxOpenConns(10)
	return &Xorm{Engine: engine}
}
