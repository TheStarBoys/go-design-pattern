package memento

import "fmt"

// 备忘录空接口
type Memento interface {}

type Originator struct {
	// 原发器的状态
	state string
}

type memo struct {
	state string
}

// 创建备忘录
func (o Originator) CreateMemento() Memento {
	fmt.Println("create memento...")
	memo := memo{state:o.state}
	return memo
}

// 加载备忘录
func (o *Originator) LoadMemento(memento Memento) {
	fmt.Println("load memento...")
	m := memento.(memo)
	o.state = m.state
}

var i int
func (o *Originator) ChangeState() {
	i %= 4
	switch i {
	case 0:
		o.state = "default state"
	case 1:
		o.state = "ready state"
	case 2:
		o.state = "run state"
	case 3:
		o.state = "end state"
	}
	i++
	o.PrintState()
}

func (o Originator) PrintState() {
	fmt.Printf("current state: %s\n", o.state)
}