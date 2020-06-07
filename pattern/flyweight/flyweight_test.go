package flyweight

import (
	"testing"
	"fmt"
)

func TestFlyweight(t *testing.T) {
	factory := FlyweightFactory{
		fsMap: map[string]Flyweight{},
	}
	f := factory.GetFlyweight("hello")
	fmt.Println(f) // &{some state}
}