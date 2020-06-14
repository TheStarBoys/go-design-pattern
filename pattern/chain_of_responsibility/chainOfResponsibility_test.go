package chain_of_responsibility

import (
	"testing"
	"fmt"
)

func TestChainOfResponsibility(t *testing.T) {
	// 先组装职责链
	h1 := new(GeneralManager)
	h2 := new(DepManager)
	h3 := new(ProjectManager)
	end := new(EndOfChain)
	h3.SetSuccessor(h2)
	h2.SetSuccessor(h1)
	h1.SetSuccessor(end)

	res1 := h3.HandleFeeRequest("小李", 300)
	fmt.Println(res1)
	res2 := h3.HandleFeeRequest("小张", 300)
	fmt.Println(res2)

	res3 := h3.HandleFeeRequest("小李", 600)
	fmt.Println(res3)
	res4 := h3.HandleFeeRequest("小张", 600)
	fmt.Println(res4)

	res5 := h3.HandleFeeRequest("小李", 1200)
	fmt.Println(res5)
	res6 := h3.HandleFeeRequest("小张", 1200)
	fmt.Println(res6)

	res7 := h3.HandleFeeRequest("小李", 12000)
	fmt.Println(res7)
	res8 := h3.HandleFeeRequest("小张", 12000)
	fmt.Println(res8)
	// Output:
	// 项目经理同意小李聚餐费用 300.00 元的请求
	// 项目经理不同意小张聚餐费用 300.00 元的请求
	// 部门经理同意小李聚餐费用 600.00 元的请求
	// 部门经理不同意小张聚餐费用 600.00 元的请求
	// 总经理同意小李聚餐费用 1200.00 元的请求
	// 总经理不同意小张聚餐费用 1200.00 元的请求
	// 小李，你小子狮子大开口啊？聚个餐还想要 12000.00 元？
	// 小张，你小子狮子大开口啊？聚个餐还想要 12000.00 元？
}