package gee

import (
	"github.com/xylong/gee/annotate"
	"reflect"
)

type BeanFactory struct {
	beans []interface{}
}

func NewBeanFactory() *BeanFactory {
	bf := &BeanFactory{
		beans: make([]interface{}, 0),
	}
	bf.beans = append(bf.beans, bf)
	return bf
}

// setBean 往内存中塞入bean
func (bf *BeanFactory) setBean(beans ...interface{}) {
	bf.beans = append(bf.beans, beans...)
}

// getBean 获取内存中预先设置的bean
func (bf *BeanFactory) getBean(p reflect.Type) interface{} {
	for _, bean := range bf.beans {
		if p == reflect.TypeOf(bean) {
			return p
		}
	}
	return nil
}

// GetBean 外部获取bean
func (bf *BeanFactory) GetBean(bean interface{}) interface{} {
	return bf.getBean(reflect.TypeOf(bean))
}

func (bf *BeanFactory) inject(controller Controller) {
	vc := reflect.ValueOf(controller).Elem()
	vct := reflect.TypeOf(controller).Elem()

	for i := 0; i < vc.NumField(); i++ {
		f := vc.Field(i)
		// 判断控制器属性是否已经实例化
		if f.Kind() != reflect.Ptr || f.IsNil() {
			continue
		}
		// 注解判断
		if annotate.IsAnnotation(f.Type()) {
			f.Set(reflect.New(f.Type().Elem()))
			f.Interface().(annotate.Annotation).SetTag(vct.Field(i).Tag)
			bf.Inject(f.Interface())
			continue
		}

		if p := bf.getBean(f.Type()); p != nil {
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}

// Inject 将bean注入到控制器
func (bf *BeanFactory) Inject(any interface{}) {
	v := reflect.ValueOf(any)

	if v.Kind() != reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		if f.Kind() != reflect.Ptr || f.IsNil() {
			continue
		}

		if p := bf.getBean(f.Type()); p != nil && f.CanInterface() {
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}
