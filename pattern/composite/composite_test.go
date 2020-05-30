package composite

import "testing"

func TestComposite(t *testing.T) {
	root := &Composite{name: "服装"}
	// 定义所有的组合对象
	c1 := &Composite{name: "男装"}
	c2 := &Composite{name: "女装"}

	// 定义所有叶子对象
	leaf1 := &Leaf{name: "衬衣"}
	leaf2 := &Leaf{name: "夹克"}
	leaf3 := &Leaf{name: "裙子"}
	leaf4 := &Leaf{name: "套装"}

	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(leaf1)
	c1.AddChild(leaf2)
	c2.AddChild(leaf3)
	c2.AddChild(leaf4)

	// 调用根对象的输出功能来输出整棵树
	root.PrintStrurt("")
}