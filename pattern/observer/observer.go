package observer

import "container/list"

// Subjecter 目标对象接口
type Subjecter interface {
	Attach(observer Observer)
	Detach(observer Observer)
	NotifyObservers()
}

// Subject 目标对象，它知道观察它的观察者，并提供注册和删除观察者的接口
type Subject struct {
	observers list.List // 用来保存观察者对象
}

// 注册观察者对象
func (s *Subject) Attach(observer Observer) {
	s.observers.PushBack(observer)
}

// 删除观察者对象
func (s *Subject) Detach(observer Observer) {
	for e := s.observers.Front(); e != nil; e = e.Next() {
		ob := e.Value.(Observer)
		if ob == observer {
			s.observers.Remove(e)
		}
	}
}

// 通知所有注册的观察者对象
func (s *Subject) NotifyObservers() {
	// 此处仅是一个示例
	// 具体的目标对象应该将组合来的NotifyObservers()方法进行改写
	for e := s.observers.Front(); e != nil; e = e.Next() {
		ob := e.Value.(Observer)
		ob.Update(s)
	}
}

// 具体的目标对象，负责把有关状态存入到相应的观察者对象
// 并在自己状态发生改变时，通知各个观察者
type ConcreteSubject struct {
	Subject
	subjectState string // 目标对象的状态
}

func (s *ConcreteSubject) GetSubjectState() string {
	return s.subjectState
}

func (s *ConcreteSubject) SetSubjectState(subjectState string) {
	s.subjectState = subjectState
	// 状态发生了改变，通知各个观察者
	s.NotifyObservers()
}

// 替换了匿名组合的同名方法，通知所有注册的观察者对象
func (s *ConcreteSubject) NotifyObservers() {
	for e := s.observers.Front(); e != nil; e = e.Next() {
		ob := e.Value.(Observer)
		ob.Update(s)
	}
}

// 观察者接口，定义一个更新的接口给那些在目标发生改变的时候被通知的对象
type Observer interface {
	// 更新的接口
	Update(subjecter Subjecter)
}

type ConcreteObserver struct {
	observerState string // 观察者状态
}

func (ob *ConcreteObserver) Update(subjecter Subjecter) {
	// 具体的更新实现
	// 这里可能需要更新观察者的状态，使其与目标的状态保持一致
	ob.observerState = subjecter.(*ConcreteSubject).GetSubjectState()
}
