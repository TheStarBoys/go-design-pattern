package decorator

import "testing"

func TestDecorator(t *testing.T) {
	old := &ConcreteComponent{}
	old.Operation()
	// Output:
	// Component Operation!

	newA := NewConcreteDecoratorA(old)
	newA.Operation()
	// Output:
	// state: default
	// Component Operation!

	newB := NewConcreteDecoratorB(old)
	newB.Operation()
	// Output:
	// Before Component Operation
	// Component Operation!

	newC := NewConcreteDecoratorB(newA)
	newC.Operation()
	// Output:
	// Before Component Operation
	// state: default
	// Component Operation!
}
