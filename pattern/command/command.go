package command

import "fmt"

// Command 命令接口，声明执行的操作
type Command interface {
	// 执行命令对应的操作
	Execute()
}

// ConcreteCommand 具体的命令实现对象
type ConcreteCommand struct {
	// 持有相应的接收者对象
	receiver *Receiver
	// 命令对象可以有自己的状态
	state string
}

func (c *ConcreteCommand) Execute() {
	// 通常会转调接收者对象的相应方法，让接收者来真正执行功能
	c.receiver.action()
}

// Receiver 接收者对象
type Receiver struct {}

// 真正执行命令相应的操作
func (r *Receiver) action() {
	// 真正执行命令操作的功能代码
	fmt.Println("action")
}

// Invoker 调用者
type Invoker struct {
	// 持有命令对象
	command Command
}

// 设置调用者持有的命令对象
func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

// 要求命令执行请求
func (i *Invoker) RunCommand() {
	// 调用命令对象的执行方法
	i.command.Execute()
}

// Client 并不是通常意义上的客户端，主要功能是要创建命令对象并设定它的接收者，因此这里并没有调用执行的代码
type Client struct {}

// 负责创建命令对象，并设定它的接收者
func (c *Client) Assemble() {
	// 创建接收者
	receiver := new(Receiver)
	// 创建命令对象，设定它的接收者
	command := &ConcreteCommand{
		receiver: receiver,
	}
	// 创建Invoker，把命令对象设置进去
	invoker := new(Invoker)
	invoker.SetCommand(command)
}