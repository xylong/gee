package main

import (
	"github.com/xylong/gee"
	v1 "github.com/xylong/gee/example/ctrl/v1"
	v2 "github.com/xylong/gee/example/ctrl/v2"
	"github.com/xylong/gee/example/middleware"
)

func main() {
	gee.Init().
		Attach(middleware.NewAuthorize()).
		Mount("v1", v1.NewUser()).
		Mount("v2", v2.NewUser()).
		Go()
}
