package prototype

import "fmt"

// OrderApi 订单的接口，声明了可以克隆自身的方法
type OrderApi interface {
	GetOrderProductNum() int
	SetOrderProductNum(num int)
	Clone() OrderApi // 克隆自身的方法
}

// PersonalOrder 个人订单对象
type PersonalOrder struct {
	CustomerName    string
	ProductId       string
	orderProductNum int
}

func (p *PersonalOrder) GetOrderProductNum() int {
	return p.orderProductNum
}

func (p *PersonalOrder) SetOrderProductNum(num int) {
	p.orderProductNum = num
}

func (p *PersonalOrder) Clone() OrderApi {
	return &PersonalOrder{
		CustomerName:    p.CustomerName,
		ProductId:       p.ProductId,
		orderProductNum: p.orderProductNum,
	}
}

func (p *PersonalOrder) String() string {
	return fmt.Sprintf("该个人订单的订购人是: %s, 订购产品是: %s, 订购数量为: %d",
		p.CustomerName, p.ProductId, p.orderProductNum)
}


// EnterpriseOrder 企业订单对象
type EnterpriseOrder struct {
	EnterpriseName  string
	ProductId       string
	orderProductNum int
}

func (p *EnterpriseOrder) GetOrderProductNum() int {
	return p.orderProductNum
}

func (p *EnterpriseOrder) SetOrderProductNum(num int) {
	p.orderProductNum = num
}

func (p *EnterpriseOrder) Clone() OrderApi {
	return &EnterpriseOrder{
		EnterpriseName:  p.EnterpriseName,
		ProductId:       p.ProductId,
		orderProductNum: p.orderProductNum,
	}
}

func (p *EnterpriseOrder) String() string {
	return fmt.Sprintf("该企业订单的订购企业是: %s, 订购产品是: %s, 订购数量为: %d",
		p.EnterpriseName, p.ProductId, p.orderProductNum)
}

// OrderBusiness 订单业务对象
type OrderBusiness struct {
	//order OrderApi 持有该需要克隆的对象
}

// 标准的原型模式不是通过传入参数得到需要克隆的对象，而是持有该需要克隆的对象
func (o *OrderBusiness) SaveOrder(order OrderApi) {
	// 1. 判断当前的预订产品数量是否大于 1000
	for order.GetOrderProductNum() > 1000 {
		// 2.1 克隆一份新订单
		newOrder := order.Clone()
		newOrder.SetOrderProductNum(1000)
		// 2.2 原来订单保留，数量-1000
		order.SetOrderProductNum(order.GetOrderProductNum() - 1000)
		// 2.3 业务功能处理，简化为打印新订单
		fmt.Println("拆分订单：", newOrder)
	}
	// 3. 不超过，直接业务处理
	fmt.Println("订单：", order)
}