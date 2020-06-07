package memento

import (
	"testing"
	"encoding/gob"
	"bytes"
	"fmt"
)

func TestMemento(t *testing.T) {
	o := Originator{}
	// 进入运行模式
	for i := 0; i < 3; i++ {
		o.ChangeState()
	}
	// 创建备忘录，接下来的操作可能会使系统结束运行
	m := o.CreateMemento()
	for i := 3; i < 5; i++ {
		o.ChangeState()
	}
	// 此时的默认状态，但期望恢复到运行状态
	// 从备忘录里加载
	o.LoadMemento(m)
	o.PrintState()
	// Output:
	// current state: default state
	// current state: ready state
	// current state: run state
	// create memento...
	// current state: end state
	// current state: default state
	// load memento...
	// current state: run state

	// 由于没有可导出字段，将无法保存到gob，json并恢复
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	encoder.Encode(m)
	fmt.Println(string(buf.Bytes()))

	tmp := memo{}
	decoder := gob.NewDecoder(bytes.NewReader(buf.Bytes()))
	decoder.Decode(&tmp)
	fmt.Println(tmp)
}
