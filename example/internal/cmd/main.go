package main

import (
	"github.com/xylong/gee"
	"github.com/xylong/gee/db"
	v1 "github.com/xylong/gee/example/internal/api/v1"
	v2 "github.com/xylong/gee/example/internal/api/v2"
	"github.com/xylong/gee/example/internal/ctrl"
	"github.com/xylong/gee/example/internal/middleware"
)

func main() {
	gee.Init().
		Beans(db.NewGorm(), db.NewXorm()).
		Attach(middleware.NewAuthorize()).
		Mount("v1", v1.NewUser()).
		Mount("v2", v2.NewUser()).
		Mount("", ctrl.NewHome()).
		Go()
}
