package mediator

import "testing"

func TestMediator(t *testing.T) {
	mediator := NewConcreteMediator()
	ca := NewConcreteColleagueA(mediator)
	cb := NewConcreteColleagueB(mediator)
	mediator.ColleagueA = ca
	mediator.ColleagueB = cb
	ca.SomeOperation()
	cb.SomeOperation()
}
