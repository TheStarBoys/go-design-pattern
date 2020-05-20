package simple_factory

import "testing"

func TestSimpleFactory(t *testing.T) {
	a := NewApi(1)
	a.Operation() // There is A

	b := NewApi(2)
	b.Operation() // There is B
}
