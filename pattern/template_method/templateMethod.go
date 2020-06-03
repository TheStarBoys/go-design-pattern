package template_method

import "fmt"

type MethodInterface interface {
	// DoPrimitiveOperation1 原语操作1，所谓原语操作就是抽象的操作，必须要由子类提供实现
	DoPrimitiveOperation1()
	// DoPrimitiveOperation2 原语操作2，所谓原语操作就是抽象的操作，必须要由子类提供实现
	DoPrimitiveOperation2()
}

// Template 定义模板方法、原语操作等的抽象类
type Template struct {
	MethodInterface
}

// 模板不需要知道这些方法的具体实现，只需要知道传入的 handle 参数实现了这些方法
func NewTemplate(handle MethodInterface) *Template {
	return &Template{
		MethodInterface: handle,
	}
}

// 模板自己再实现接口，显式的panic，强制子类去实现，而不是子类不实现，导致让人迷惑的栈溢出的错误
// 可省略，但子类不实现的话，报的错让人迷惑
func (t *Template) DoPrimitiveOperation1() {
	panic("need to implement")
}

func (t *Template) DoPrimitiveOperation2() {
	panic("need to implement")
}

// TemplateMethod 模板方法，定义算法骨架
func (t *Template) TemplateMethod() {
	// 显式的调用接口的方法，而不是调用模板自己重写的方法
	t.MethodInterface.DoPrimitiveOperation1()
	t.MethodInterface.DoPrimitiveOperation2()
}

// Concrete 具体实现类，实现原语操作
type Concrete struct {
	*Template
}

func NewConcrete() *Concrete {
	c := &Concrete{}
	c.Template = NewTemplate(c)

	return c
}

// 如果不提供具体的实现，将导致栈溢出
func (c *Concrete) DoPrimitiveOperation1() {
	// 具体的实现
	fmt.Println("DoPrimitiveOperation1")
}

func (c *Concrete) DoPrimitiveOperation2() {
	// 具体的实现
	fmt.Println("DoPrimitiveOperation2")
}