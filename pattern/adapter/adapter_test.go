package adapter

import "testing"

func TestAdapter(t *testing.T) {
	adaptee := new(AdapteeImpl)
	target := NewAdapter(adaptee)
	target.Request() // specific request
}