package command1

import "fmt"

// MainBoardApi 主板的接口
type MainBoardApi interface {
	// 主板具有能开机的功能
	Open()
}

// GigaMainBoard 技嘉主板，开机命令的真正实现者，在Command模式中充当Receiver
type GigaMainBoard struct {}

// 真正的开机命令的实现
func (g *GigaMainBoard) Open() {
	fmt.Println("技嘉主板现在正在开机，请稍等......")
	fmt.Println("接通电源......")
	fmt.Println("设备检查......")
	fmt.Println("装载系统......")
	fmt.Println("机器正常运转起来......")
	fmt.Println("机器已经正常打开，请操作......")
}

// MSIMainBoard 微星主板，开机命令的真正实现者，在Command模式中充当Receiver
type MSIMainBoard struct {}

// 真正的开机命令的实现
func (m *MSIMainBoard) Open() {
	fmt.Println("微星主板现在正在开机，请稍等......")
	fmt.Println("接通电源......")
	fmt.Println("设备检查......")
	fmt.Println("装载系统......")
	fmt.Println("机器正常运转起来......")
	fmt.Println("机器已经正常打开，请操作......")
}

// Command 命令接口，声明执行的操作
type Command interface {
	// 执行命令对应的操作
	Execute()
}

// OpenCommand 开机命令的实现
type OpenCommand struct {
	// 持有真正实现命令的接收者——主板对象
	mainBoard MainBoardApi
}

func NewOpenCommand(mainBoard MainBoardApi) *OpenCommand {
	return &OpenCommand{
		mainBoard: mainBoard,
	}
}

func (o *OpenCommand) Execute() {
	// 对于命令对象，根本不知道如何开机，会转调主板对象
	// 让主板去完成开机功能
	o.mainBoard.Open()
}

// Box 机箱对象，本身有按钮，持有按钮对应的命令对象
type Box struct {
	// 开机命令对象
	OpenCommand Command
}

// 设置开机命令对象
func (b *Box) SetOpenCommand(command Command) {
	b.OpenCommand = command
}

// 提供给客户使用，接收并相应用户请求，相当于按钮被按下触发的方法
func (b *Box) OpenButtonPressed() {
	// 按下按钮，执行命令
	b.OpenCommand.Execute()
}