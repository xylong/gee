基于gin的脚手架

### Installation
安装脚手架
```
go get -u github.com/xylong/gee
```
### 快速开始
```
package main

import (
	v1 "blog/ctrl/v1"
	"github.com/xylong/gee"
)

func main() {
	gee.Init().Mount("v1",v1.NewUser()).Go()
}

```