package visitor

import (
	"testing"
)

func TestVisitor(t *testing.T) {
	os := new(ObjectStructure)
	// 准备客户
	cm1 := &EnterpriseCustomer {
		BasicCustomer: BasicCustomer{name:"ABC 集团"},
	}
	cm2 := &EnterpriseCustomer {
		BasicCustomer: BasicCustomer{name:"CDE 公司"},
	}
	cm3 := &PersonalCustomer{
		BasicCustomer: BasicCustomer{name:"张三"},
	}
	// 添加到对象结构
	os.AddElement(cm1)
	os.AddElement(cm2)
	os.AddElement(cm3)
	// 客户提出服务请求
	os.HandleRequest(new(ServiceRequestVisitor))
	// 对客户进行偏好分析
	os.HandleRequest(new(PredilectionAnalyzeVisitor))
	// Output:
	// ABC 集团企业提出服务请求
	// CDE 公司企业提出服务请求
	// 客户张三提出服务请求
	// 现在对企业客户ABC 集团进行产品偏好分析
	// 现在对企业客户CDE 公司进行产品偏好分析
	// 现在对个人客户张三进行产品偏好分析
}