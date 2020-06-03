package strategy

import "fmt"

// Strategy 策略，定义算法的接口
type Strategy interface {
	// 某个算法接口
	algorithmInterface()
}

// ConcreteStrategyA 实现具体的算法
type ConcreteStrategyA struct {}

func (c ConcreteStrategyA) algorithmInterface() {
	// 具体的算法实现
	fmt.Println("StrategyA")
}

// ConcreteStrategyB 实现具体的算法
type ConcreteStrategyB struct {}

func (c ConcreteStrategyB) algorithmInterface() {
	// 具体的算法实现
	fmt.Println("StrategyB")
}

// ConcreteStrategyC 实现具体的算法
type ConcreteStrategyC struct {}

func (c ConcreteStrategyC) algorithmInterface() {
	// 具体的算法实现
	fmt.Println("StrategyC")
}

// Context 上下文对象
type Context struct {
	// 持有策略
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{
		strategy:strategy,
	}
}

// 上下文为客户提供的操作接口
func (c *Context) contextInterface() {
	// 通常用转调具体的策略的具体算法
	c.strategy.algorithmInterface()
}