package simple_factory

import "fmt"

// 接口定义，该接口可以通过简单工厂来创建
type Api interface {
	Operation()
}

// ImplA 接口的具体实现对象A
type ImplA struct {}

func (a *ImplA) Operation() {
	fmt.Println("There is A")
}

// ImplB 接口的具体实现对象B
type ImplB struct {}

func (b *ImplB) Operation() {
	fmt.Println("There is B")
}

// NewApi 根据condition来选择具体的实现
// 如果只有一个实现，可以省略条件，因为没有选择的必要
func NewApi(condition int) Api {
	switch condition {
	case 1:
		return &ImplA{}
	case 2:
		return &ImplB{}
	}

	return nil
}