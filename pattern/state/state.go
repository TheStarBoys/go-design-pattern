package state

import "fmt"

// 封装与Context 的一个特定状态相关的行为
type State interface {
	Handle(parameter string)
}

type ConcreteStateA struct {}

func (c *ConcreteStateA) Handle(parameter string) {
	// 实现具体的处理
	fmt.Println("ConcreteStateA:", parameter)
}

type ConcreteStateB struct {}

func (c *ConcreteStateB) Handle(parameter string) {
	// 实现具体的处理
	fmt.Println("ConcreteStateB:", parameter)
}

// 定义用户感兴趣的方法
type Context struct {
	// 持有一个状态
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

// 用户感兴趣的方法
func (c *Context) Request(parameter string) {
	// 在处理中，会转调state来处理
	c.state.Handle(parameter)
}