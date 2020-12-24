package main

import (
	"github.com/xylong/gee"
	"github.com/xylong/gee/db"
	v1 "github.com/xylong/gee/example/api/v1"
	v2 "github.com/xylong/gee/example/api/v2"
	"github.com/xylong/gee/example/middleware"
)

func main() {
	gee.Init().
		Orm(db.NewGorm(), db.NewXorm()).
		Attach(middleware.NewAuthorize()).
		Mount("v1", v1.NewUser()).
		Mount("v2", v2.NewUser()).
		Go()
}
