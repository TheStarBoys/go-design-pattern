package command1

import (
	"testing"
	"fmt"
)

func TestCommand(t *testing.T) {
	// 1. 把命令和真正的实现组合起来，相当于在组装机器
	mainBoard := new(GigaMainBoard)
	// 创建开机命令
	openCommand := NewOpenCommand(mainBoard)
	// 创建重启命令
	resetCommand := NewResetCommand(mainBoard)
	// 2. 为机箱上的按钮设置对应的命令，让按钮知道该干什么
	box := new(Box)
	// 先正确配置，也就是开机按钮对应开机命令，重启按钮对应重启命令
	box.SetOpenCommand(openCommand)
	box.SetResetCommand(resetCommand)
	// 3. 模拟按下机箱上的按钮
	fmt.Println("正确配置下--------------------------->")
	fmt.Println(">>>按下开机按钮：>>>")
	box.OpenButtonPressed()
	fmt.Println(">>>按下重启按钮：>>>")
	box.ResetButtonPressed()

	// 然后来错误配置一回
	box.SetOpenCommand(resetCommand)
	box.SetResetCommand(openCommand)
	// 4. 再来模拟按下按钮的操作
	fmt.Println("错误配置下--------------------------->")
	fmt.Println(">>>按下开机按钮：>>>")
	box.OpenButtonPressed()
	fmt.Println(">>>按下重启按钮：>>>")
	box.ResetButtonPressed()
}
