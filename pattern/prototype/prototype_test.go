package prototype

import (
	"testing"
	"fmt"
)

func TestPrototype(t *testing.T) {
	o1 := &PersonalOrder{
		CustomerName: "张三",
		ProductId: "TX001",
	}
	o1.SetOrderProductNum(2125)

	o2 := &EnterpriseOrder{
		EnterpriseName: "Excel",
		ProductId: "TX002",
	}
	o2.SetOrderProductNum(1024)

	business := OrderBusiness{}
	business.SaveOrder(o1)
	fmt.Println()
	business.SaveOrder(o2)
}
