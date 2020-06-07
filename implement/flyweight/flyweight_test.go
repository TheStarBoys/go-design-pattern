package flyweight

import (
	"testing"
	"fmt"
)

func TestFlyweight(t *testing.T) {
	manager := GetManager()
	manager.Login("张三")
	manager.Login("李四")
	b1 := manager.HasPermit("张三", "薪资数据", "查看")
	b2 := manager.HasPermit("李四", "薪资数据", "查看")

	fmt.Printf("b1: %v, b2: %v\n", b1, b2)
	// Output:
	// fm: 0xc000004520
	// fm: 0xc000004520
	// fm: 0xc000004560
	// b1: false, b2: true
}
