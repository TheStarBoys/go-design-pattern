package facade

import "fmt"

// A模块接口
type AModuleApi interface {
	TestA()
}

// A模块实现
type AModuleImpl struct {}

func (a *AModuleImpl) TestA() {
	fmt.Println("A module")
}

// B模块接口
type BModuleApi interface {
	TestB()
}

// B模块实现
type BModuleImpl struct {}

func (b *BModuleImpl) TestB() {
	fmt.Println("B module")
}

// C模块接口
type CModuleApi interface {
	TestC()
}

// C模块实现
type CModuleImpl struct {}

func (c *CModuleImpl) TestC() {
	fmt.Println("C module")
}

// 外观对象
type Facade struct {}

// 对于客户想要使用的功能，提供了统一的外观
// 客户不需要关心内部调用了哪些模块，只需要关心最终是否能达到效果
func (f *Facade) Test() {
	a := &AModuleImpl{}
	b := &BModuleImpl{}
	c := &CModuleImpl{}

	a.TestA()
	b.TestB()
	c.TestC()
}