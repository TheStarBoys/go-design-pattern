package composite

import (
	"container/list"
	"fmt"
)

// Component 组件接口
type ComponentInterface interface {
	// 组件对象可能有的功能方法
	PrintStrurt(preStr string)
	// 向组合对象中加入组合对象
	AddChild(child ComponentInterface)
	// 从组合对象中移除组合对象
	RemoveChild(child ComponentInterface)
	// 返回某个索引对应的组件对象
	GetChildren(index int) ComponentInterface
}

// 提供一个缺省的实现
type Component struct {}

func (c *Component) PrintStrurt(preStr string) {
	panic("need to implement")
}

func (c *Component) AddChild(child ComponentInterface) {
	panic("need to implement")
}

func (c *Component) RemoveChild(child ComponentInterface) {
	panic("need to implement")
}

func (c *Component) GetChildren(index int) ComponentInterface {
	panic("need to implement")
}

// Composite 组合对象，通常需要存储子对象，定义有子部件的部件行为
// 并实现在 Component 里面定义的与子部件有关的操作
type Composite struct {
	Component
	// 用来存储组合对象中包含的子组件对象
	children list.List
	name string
}

// 通常在里面需要实现递归的调用
func (c *Composite) PrintStrurt(preStr string) {
	// 先把自己输出出去
	fmt.Println(preStr + "+" + c.name)
	// 缩进
	preStr += " "
	// 输出子组件
	for e := c.children.Front(); e != nil; e = e.Next() {
		e.Value.(ComponentInterface).PrintStrurt(preStr)
	}
}

func (c *Composite) AddChild(child ComponentInterface) {
	// list包自己内部实现了 lazy init
	c.children.PushBack(child)
}

func (c *Composite) RemoveChild(child ComponentInterface) {
	for e := c.children.Front(); e != nil; e = e.Next() {
		if e.Value.(ComponentInterface) == child {
			c.children.Remove(e)
		}
	}
}

func (c *Composite) GetChildren(index int) ComponentInterface {
	e := c.children.Front()
	for ; index >= 0 && e != nil; e, index = e.Next(), index + 1 {}

	return e.Value.(ComponentInterface)
}

// Leaf 叶子对象，叶子对象不再包含其他子对象
type Leaf struct {
	Composite
	name string
}

func (l *Leaf) PrintStrurt(preStr string) {
	fmt.Println(preStr + "-" + l.name)
}