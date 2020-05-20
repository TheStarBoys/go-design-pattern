package builder

import (
	"testing"
)

func TestBuilder1(t *testing.T) {
	builder := &Builder1{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res != "123" {
		t.Errorf("builder.GetResult() err, expect: 123, actual: %s\n", res)
	}
}

func TestBuilder2(t *testing.T) {
	builder := &Builder2{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res != 6 {
		t.Errorf("builder.GetResult() err, expect: 6, actual: %d\n", res)
	}
}