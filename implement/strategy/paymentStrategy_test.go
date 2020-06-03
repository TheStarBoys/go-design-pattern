package strategy

import "testing"

func TestPaymentStrategy(t *testing.T) {
	// 创建相应策略
	RMB := new(RMBCash)
	Dollar := new(DollarCash)
	Card := new(Card)

	// 准备小李的支付工资上下文
	ctx1 := NewPaymentContext("小李", 5000,"5000108", RMB)
	// 向小李支付工资
	ctx1.PayNow()

	// 准备Petter的支付工资上下文
	ctx2 := NewPaymentContext("Petter", 8000, "5000109", Dollar)
	// 向Petter支付工资
	ctx2.PayNow()

	// 准备小张的支付工资上下文
	ctx3 := NewPaymentContext("小张", 10000, "5000110", Card)
	// 向Petter支付工资
	ctx3.PayNow()
	// Output:
	// 现在给小李人民币现金支付5000.00元
	// 现在给Petter美元现金支付8000.00元
	// 现在给小张的账号5000110支付了10000.00元
}
