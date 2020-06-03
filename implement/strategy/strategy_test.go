package strategy

import (
	"testing"
	"fmt"
)

func TestStrategy(t *testing.T) {
	// 1. 选择并创建需要使用的策略对象
	var strategy Strategy
	strategy = new(LargeCustomerStrategy)
	// 2. 创建上下文
	price := NewPrice(strategy)
	// 3. 计算报价
	quote := price.Quote(1000)
	fmt.Println("向客户报价：", quote)
	// Output:
	// 对于大客户，统一折扣 10%
	// 向客户报价： 900
}
