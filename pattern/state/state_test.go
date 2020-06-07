package state

import "testing"

func TestState(t *testing.T) {
	var state State
	state = new(ConcreteStateA)
	ctx := new(Context)
	ctx.SetState(state)
	ctx.Request("hello")

	state = new(ConcreteStateB)
	ctx.SetState(state)
	ctx.Request("hello")
	// Output:
	// ConcreteStateA: hello
	// ConcreteStateB: hello
}