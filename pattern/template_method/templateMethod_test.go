package template_method

import "testing"

func TestTemplateMethod(t *testing.T) {
	// panic 接口必须传值
	//template := NewTemplate(nil)
	//template.TemplateMethod()

	concrete := NewConcrete()
	concrete.TemplateMethod()
}
