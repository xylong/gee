package gee

// Controller 控制器
type Controller interface {
	// Build 创建控制器
	Build(gee *Gee)
}
