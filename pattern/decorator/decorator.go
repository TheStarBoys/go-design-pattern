package decorator

import "fmt"

// 组件接口，可以给这些组件动态添加职责或状态
type Component interface {
	Operation()
}

// 实现组件接口的对象
type ConcreteComponent struct {}

func (c *ConcreteComponent) Operation() {
	// do some operation
	fmt.Println("Component Operation!")
}

// Decorator 装饰器，在 Go 中是不必要的
// type Decorator struct {
//	 component Component
// }

// 编译期检测装饰器的实现是否满足为一个组件
var _ = Component(&ConcreteDecoratorA{})
var _ = Component(&ConcreteDecoratorB{})

// 装饰器的具体实现，向组件添加职责
type ConcreteDecoratorA struct {
	Component
	state string // 新增的状态
}

func NewConcreteDecoratorA(component Component) *ConcreteDecoratorA {
	return &ConcreteDecoratorA{
		Component: component,
		state: "default",
	}
}

func (a *ConcreteDecoratorA) Operation() {
	// 调用组件的方法，可以在调用前后执行一些附加操作
	// 可以使用添加的状态
	fmt.Println("state:", a.state)
	a.Component.Operation()
}

type ConcreteDecoratorB struct {
	Component
}

func NewConcreteDecoratorB(component Component) *ConcreteDecoratorB {
	return &ConcreteDecoratorB{
		Component: component,
	}
}

func (a *ConcreteDecoratorB) addedBehavior() {
	// 添加新的职责
	fmt.Println("Before Component Operation")
}

func (a *ConcreteDecoratorB) Operation() {
	// 调用组件的方法，可以在调用前后执行一些附加操作
	a.addedBehavior()
	a.Component.Operation()
}