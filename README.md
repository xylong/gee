基于gin的脚手架

### Installation
安装脚手架
```
go get -u github.com/xylong/gee
```
### 快速开始
```go
package main

import (
	v1 "blog/ctrl/v1"
	"github.com/xylong/gee"
)

func main() {
	gee.Init().
		Orm(db.NewGorm(), db.NewXorm()).
		Attach(middleware.NewAuthorize()).
		Mount("v1", v1.NewUser()).
		Mount("v2", v2.NewUser()).
		Go()
}
```