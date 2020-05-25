package prototype

import (
	"testing"
)

func TestPrototypeManager(t *testing.T) {
	manager := NewPrototypeManager()
	t1 := &T1{name: "type1"}
	t2 := &T2{age: 3}
	manager.Set("type1", t1)
	manager.Set("type2", t2)

	tt1 := manager.Get("type1").Clone()
	tt2 := manager.Get("type2").Clone()
	if tt1 == t1 {
		t.Fatal("it should not equal")
	}
	if tt2 == t2 {
		t.Fatal("it should not equal")
	}
}

type T1 struct {
	name string
}

func (t *T1) Clone() Cloneable {
	return &T1{
		name: t.name,
	}
}

type T2 struct {
	age int
}

func (t *T2) Clone() Cloneable {
	return &T2{
		age: t.age,
	}
}
