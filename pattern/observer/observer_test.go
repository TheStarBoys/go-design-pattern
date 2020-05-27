package observer

import (
	"testing"
	"fmt"
)

func TestObserver(t *testing.T) {
	sub := new(ConcreteSubject)

	ob1 := new(ConcreteObserver)
	ob2 := new(ConcreteObserver)
	fmt.Printf("ob1: %#v, ob2: %#v\n", ob1, ob2)
	// ob1: &observer.ConcreteObserver{observerState:""}, ob2: &observer.ConcreteObserver{observerState:""}

	sub.Attach(ob1)
	sub.Attach(ob2)

	sub.SetSubjectState("hello")
	fmt.Printf("ob1: %#v, ob2: %#v\n", ob1, ob2)
	// ob1: &observer.ConcreteObserver{observerState:"hello"}, ob2: &observer.ConcreteObserver{observerState:"hello"}
}