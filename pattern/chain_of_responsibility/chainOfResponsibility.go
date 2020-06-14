package chain_of_responsibility

import (
	"fmt"
	"reflect"
)

// Handler 定义职责对象的接口
type Handler interface {
	SetSuccessor(successor Handler) // 设置下一个处理请求的对象
	HandleFeeRequest(user string, fee float64) string // 处理聚餐费用的申请
}

// BasicHandler 所有具体的实现应该继承它
type BasicHandler struct {
	successor Handler
}

func (h *BasicHandler) SetSuccessor(successor Handler) {
	h.successor = successor
}

func (h *BasicHandler) HandleFeeRequest(user string, fee float64) string {
	panic(fmt.Sprintf("%s.HandleFeeRequest() needs to be implement", reflect.TypeOf(*h).Name()))
}

// ProjectManager 项目经理
type ProjectManager struct {
	BasicHandler
}

// 建议通过该方法在编译期就确保实现了Handler接口，在Handler接口发生变化时，能及时定位
var _ = Handler(&ProjectManager{})

func (h *ProjectManager) HandleFeeRequest(user string, fee float64) string {
	str := ""
	// 项目经理的权限比较小，只能在500以内
	if fee < 500 {
		// 为了测试，只同意小李的
		if user == "小李" {
			str = fmt.Sprintf("项目经理同意%s聚餐费用 %.2f 元的请求", user, fee)
		} else {
			str = fmt.Sprintf("项目经理不同意%s聚餐费用 %.2f 元的请求", user, fee)
		}
	} else {
		// 如果超过其权限，且有后继，交给级别更高的人处理
		if h.successor != nil {
			return h.successor.HandleFeeRequest(user, fee)
		}
	}
	return str
}

// DepManager 部门经理
type DepManager struct {
	BasicHandler
}

var _ Handler = &DepManager{}

func (h *DepManager) HandleFeeRequest(user string, fee float64) string {
	str := ""
	// 部门经理的权限只能在1000以内
	if fee < 1000 {
		// 为了测试，只同意小李的
		if user == "小李" {
			str = fmt.Sprintf("部门经理同意%s聚餐费用 %.2f 元的请求", user, fee)
		} else {
			str = fmt.Sprintf("部门经理不同意%s聚餐费用 %.2f 元的请求", user, fee)
		}
	} else {
		// 如果超过其权限，且有后继，交给级别更高的人处理
		if h.successor != nil {
			return h.successor.HandleFeeRequest(user, fee)
		}
	}
	return str
}

// GeneralManager 总经理
type GeneralManager struct {
	BasicHandler
}

var _ = Handler(&GeneralManager{})

func (h *GeneralManager) HandleFeeRequest(user string, fee float64) string {
	str := ""
	// 总经理的权限很大，1000~10000的金额都可以交由他处理
	if 10000 >= fee && fee >= 1000 {
		// 为了测试，只同意小李的
		if user == "小李" {
			str = fmt.Sprintf("总经理同意%s聚餐费用 %.2f 元的请求", user, fee)
		} else {
			str = fmt.Sprintf("总经理不同意%s聚餐费用 %.2f 元的请求", user, fee)
		}
	} else {
		// 如果超过其权限，且有后继，交给级别更高的人处理
		if h.successor != nil {
			return h.successor.HandleFeeRequest(user, fee)
		}
	}
	return str
}

// EndOfChain 职责链的终点，因为可能一个请求，没有任何一个人能够处理它
type EndOfChain struct {
	BasicHandler
}

func (h *EndOfChain) HandleFeeRequest(user string, fee float64) string {
	return fmt.Sprintf("%s，你小子狮子大开口啊？聚个餐还想要 %.2f 元？", user, fee)
}