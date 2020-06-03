package strategy

import "fmt"

// Strategy 策略，定义计算报价算法的接口
type Strategy interface {
	// 计算应报价格, goodsPrice 商品销售原价
	CalcPrice(goodsPrice float64) float64
}

// 具体算法实现，为新客户或者是普通客户计算应报的价格
type NormalCustomerStrategy struct {}

func (s NormalCustomerStrategy) CalcPrice(goodsPrice float64) float64 {
	fmt.Println("对于新客户或者是普通客户，没有折扣")
	return goodsPrice
}

// 具体算法实现，为老客户计算应报的价格
type OldCustomerStrategy struct {}

func (s OldCustomerStrategy) CalcPrice(goodsPrice float64) float64 {
	fmt.Println("对于老客户，统一折扣 5%")
	return goodsPrice * (1 - 0.05)
}

// 具体算法实现，为大客户计算应报的价格
type LargeCustomerStrategy struct {}

func (s LargeCustomerStrategy) CalcPrice(goodsPrice float64) float64 {
	fmt.Println("对于大客户，统一折扣 10%")
	return goodsPrice * (1 - 0.1)
}

// Price 价格管理， 主要完成计算向客户所报价格的功能
type Price struct {
	// 持有一个策略
	strategy Strategy
}

func NewPrice(strategy Strategy) *Price {
	return &Price{
		strategy: strategy,
	}
}

// 报价，计算对客户的报价
func (p *Price) Quote(goodsPrice float64) float64 {
	return p.strategy.CalcPrice(goodsPrice)
}