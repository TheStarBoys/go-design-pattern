package iterator

import (
	"testing"
	"fmt"
)

func TestIterator(t *testing.T) {
	names := []string{"张三", "李四", "王五"}
	// 创建聚合对象
	aggregate := NewConcreteAggregate(names)
	// 循环输出聚合对象中的值
	iterator := aggregate.CreateIterator()
	for iterator.First(); iterator.IsEnd() == false; iterator.Next() {
		// 取出当前元素
		curr := iterator.CurrentItem().(string)
		fmt.Printf("currElem: %s\n", curr)
	}
}
