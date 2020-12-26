package annotate

import "reflect"

// Annotations 注解列表
var Annotations []Annotation

// Annotation 注解
type Annotation interface {
	SetTag(tag reflect.StructTag)
}

func init() {
	Annotations = make([]Annotation, 0)
	Annotations = append(Annotations, new(Value))
}

// IsAnnotation 判断当前注入对象是否是注解
func IsAnnotation(p reflect.Type) bool {
	for _, item := range Annotations {
		if reflect.TypeOf(item) == p {
			return true
		}
	}
	return false
}
