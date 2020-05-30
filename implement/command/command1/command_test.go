package command1

import "testing"

func TestCommand(t *testing.T) {
	// 1. 把命令和真正的实现组合起来，相当于在组装机器
	mainBoard := new(GigaMainBoard)
	// 把机箱上按钮的连接线插到主板上
	openCommand := NewOpenCommand(mainBoard)
	// 2. 为机箱上的按钮设置对应的命令，让按钮知道该干什么
	box := new(Box)
	box.SetOpenCommand(openCommand)
	// 3. 模拟按下机箱上的按钮
	box.OpenButtonPressed()
}
