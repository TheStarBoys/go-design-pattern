package adapter

import "fmt"

// Target 是适配的目标接口
type Target interface {
	Request() // 客户端请求处理的方法
}

// Adaptee 已经存在的接口，这个接口需要被适配
type Adaptee interface {
	SpecificRequest()
}

// AdapteeImpl 需要被适配的接口的实现
type AdapteeImpl struct {}

// SpecificRequest 原本已经存在的，特殊请求的方法
func (a *AdapteeImpl) SpecificRequest() {
	fmt.Println("specific request")
}

// Adapter 适配器对象，将原本不适配的接口进行适配
type Adapter struct {
	adaptee Adaptee
}

// NewAdapter 是Adapter的工厂函数
func NewAdapter(adaptee Adaptee) Target {
	return &Adapter{
		adaptee: adaptee,
	}
}

// Request 实现Target接口
func (a *Adapter) Request() {
	a.adaptee.SpecificRequest()
}