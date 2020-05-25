package mediator

// Mediator 中介者，定义各个同事对象通信的接口
type Mediator interface {
	// 同事对象在自身改变的时候来通知中介者的方法
	// 让中介者去负责相应的与其他同事对象的交互
	// colleague 同事对象自身，好让中介者对象通过对象实例去获取同事对象的状态
	Changed(colleague ColleagueInterface)
}

// 中介者实现
type ConcreteMediator struct {
	// 持有并维护同事A
	ColleagueA *ConcreteColleagueA
	// 持有并维护同事B
	ColleagueB *ConcreteColleagueB
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{}
}

func (m *ConcreteMediator) Changed(colleague ColleagueInterface) {
	// 某个同事发生了变化，通常需要与其他同事交互
	// 具体协调相应的同事对象来实现协作行为
	switch colleague.(type) {
	case *ConcreteColleagueA:
		// 进行一些处理
	case *ConcreteColleagueB:
		// 进行一些处理
	}
}

// Colleague 同事接口
type ColleagueInterface interface {
	GetMediator() Mediator // 获取当前同事类对应的中介者对象
}

// 同事基类，实现了同事接口，同事的真正实现将匿名组合该类
type Colleague struct {
	mediator Mediator
}

func (c *Colleague) GetMediator() Mediator {
	return c.mediator
}

// ConcreteColleagueA 同事A
type ConcreteColleagueA struct {
	Colleague
}

func NewConcreteColleagueA(mediator Mediator) *ConcreteColleagueA {
	return &ConcreteColleagueA {
		Colleague: Colleague {
			mediator:mediator,
		},
	}
}

// SomeOperation 执行某些业务功能
func (c *ConcreteColleagueA) SomeOperation() {
	// 在需要跟其他同事通信的时候，通知中介者对象
	c.GetMediator().Changed(c)
}

// ConcreteColleagueB 同事B
type ConcreteColleagueB struct {
	Colleague
}

func NewConcreteColleagueB(mediator Mediator) *ConcreteColleagueB {
	return &ConcreteColleagueB {
		Colleague: Colleague {
			mediator:mediator,
		},
	}
}

// SomeOperation 执行某些业务功能
func (c *ConcreteColleagueB) SomeOperation() {
	// 在需要跟其他同事通信的时候，通知中介者对象
	c.GetMediator().Changed(c)
}