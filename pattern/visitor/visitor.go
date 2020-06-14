package visitor

import (
	"fmt"
)

type Customer interface {
	// 接受访问者的访问
	Accept(visitor Visitor)
}

type BasicCustomer struct {
	name string
	customerId string
}

// EnterpriseCustomer 企业客户
type EnterpriseCustomer struct {
	BasicCustomer
	linkman string
	linkTelephone string
	registerAddress string
}

func (c *EnterpriseCustomer) Accept(visitor Visitor) {
	// 回调访问者相应的方法
	visitor.VisitEnterpriseCustomer(c)
}

// PersonalCustomer 个人客户
type PersonalCustomer struct {
	BasicCustomer
	telephone string
	age int
}

func (c *PersonalCustomer) Accept(visitor Visitor) {
	// 回调访问者相应的方法
	visitor.VisitPersonalCustomer(c)
}

// Visitor 访问者接口
type Visitor interface {
	// 访问企业客户，相当于给企业客户添加访问者的功能
	VisitEnterpriseCustomer(ec *EnterpriseCustomer)
	// 访问个人客户，相当于给个人客户添加访问者的功能
	VisitPersonalCustomer(pc *PersonalCustomer)
}

// ServiceRequestVisitor 具体的访问者，实现客户提出服务请求的功能
type ServiceRequestVisitor struct {}

func (v *ServiceRequestVisitor) VisitEnterpriseCustomer(ec *EnterpriseCustomer) {
	fmt.Printf("%s企业提出服务请求\n", ec.name)
}

func (v *ServiceRequestVisitor) VisitPersonalCustomer(pc *PersonalCustomer) {
	fmt.Printf("客户%s提出服务请求\n", pc.name)
}

type PredilectionAnalyzeVisitor struct {

}

func (v *PredilectionAnalyzeVisitor) VisitEnterpriseCustomer(ec *EnterpriseCustomer) {
	// 根据以往购买的历史、潜在购买意向
	// 以及客户所在行业的发展趋势、客户的发展预期等的分析
	fmt.Printf("现在对企业客户%s进行产品偏好分析\n", ec.name)
}

func (v *PredilectionAnalyzeVisitor) VisitPersonalCustomer(pc *PersonalCustomer) {
	fmt.Printf("现在对个人客户%s进行产品偏好分析\n", pc.name)
}

type ObjectStructure struct {
	col []Customer // 要操作的客户集合
}

// 提供给客户端操作的高层接口，具体的功能由客户端传入的访问者决定
func (o *ObjectStructure) HandleRequest(visitor Visitor) {
	for _, customer := range o.col {
		customer.Accept(visitor)
	}
}

// 向对象结构中添加元素
func (o *ObjectStructure) AddElement(elem Customer) {
	o.col = append(o.col, elem)
}