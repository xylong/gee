package annotate

import "reflect"

type Value struct {
	tag reflect.StructTag
}

func (v *Value) SetTag(tag reflect.StructTag) {
	v.tag = tag
}

func (v *Value) String() string {
	return "18"
}
