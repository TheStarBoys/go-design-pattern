package strategy

import "fmt"

// 支付工资的策略接口，公司有多种支付工资的算法
// 比如：现金、银行卡、现金加股票、现金加期权、美元支付等
type PaymentStrategy interface {
	// 公司给某人真正支付工资, ctx 支付工资的上下文，里面包含算法需要的数据
	Pay(ctx *PaymentContext)
}

// 人民币现金支付
type RMBCash struct {}

func (c *RMBCash) Pay(ctx *PaymentContext) {
	fmt.Printf("现在给%s人民币现金支付%.2f元\n",
		ctx.GetUserName(), ctx.GetMoney())
}

// 美元现金支付
type DollarCash struct {}

func (c DollarCash) Pay(ctx *PaymentContext) {
	fmt.Printf("现在给%s美元现金支付%.2f元\n",
		ctx.GetUserName(), ctx.GetMoney())
}

// 新的策略的实现，支付到银行卡
type Card struct {}

func (c Card) Pay(ctx *PaymentContext) {
	fmt.Printf("现在给%s的账号%s支付了%.2f元\n",
		ctx.GetUserName(), ctx.GetAccount(), ctx.GetMoney())
}

// 支付工资的上下文，每个人的工资不同，支付方式也不同
type PaymentContext struct {
	// 应被支付工资的人员
	userName string
	// 员工的银行账户
	account string
	// 应被支付的工资金额
	money float64
	// 持有支付工资的方式的策略接口
	strategy PaymentStrategy
}

func NewPaymentContext(userName string, money float64, account string, strategy PaymentStrategy) *PaymentContext {
	return &PaymentContext{
		userName: userName,
		money: money,
		account: account,
		strategy: strategy,
	}
}

func (ctx PaymentContext) GetUserName() string {
	return ctx.userName
}

func (ctx PaymentContext) GetMoney() float64 {
	return ctx.money
}

func (ctx PaymentContext) GetAccount() string {
	return ctx.account
}

// 立即支付工资
func (ctx *PaymentContext) PayNow() {
	// 使用客户希望的支付策略来支付工资
	ctx.strategy.Pay(ctx)
}