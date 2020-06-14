package bridge

import "fmt"

// MessageImplementor 实现发送消息的统一接口
type MessageImplementor interface {
	Send(message, toUser string)
}

// 利用go接口可以接收值的特性，根本不需要通过 AbstractMessage 来操作

/*
// AbstractMessage 抽象的消息对象
type AbstractMessage struct {
	impl MessageImplementor // 持有一个消息接口的实现
}

func (a *AbstractMessage) SendMessage(message, toUser string) {
	// 转调具体的实现
	a.impl.Send(message, toUser)
}
*/

// MessageSMS 站内短消息的方式发送消息
type MessageSMS struct {}

func (m *MessageSMS) Send(message, toUser string) {
	fmt.Printf("使用站内短信方式，发送消息'%s'给%s\n", message, toUser)
}

// MessageEmail E-mail的方式发送消息
type MessageEmail struct {}

func (m *MessageEmail) Send(message, toUser string) {
	fmt.Printf("使用E-mail方式，发送消息'%s'给%s\n", message, toUser)
}

// MessageMobile 以手机短信的方式发送消息
type MessageMobile struct {}

func (m *MessageMobile) Send(message, toUser string) {
	fmt.Printf("使用手机短信方式，发送消息'%s'给%s\n", message, toUser)
}

// CommonMessage 普通消息
type CommonMessage struct {
	//AbstractMessage // 什么也不干，直接就是AbstractMessage的方法调用过程
	MessageImplementor
}

func NewCommonMessage(impl MessageImplementor) *CommonMessage {
	return &CommonMessage{
		MessageImplementor: impl,
	}
}

// UrgencyMessage 加急消息
type UrgencyMessage struct {
	MessageImplementor
}

func NewUrgencyMessage(impl MessageImplementor) *UrgencyMessage {
	return &UrgencyMessage{
		MessageImplementor: impl,
	}
}

func (m *UrgencyMessage) SendMessage(message, toUser string) {
	message += "加急："
	m.MessageImplementor.Send(message, toUser)
}

// 扩展自己的新功能：监控某消息的处理过程
func (m *UrgencyMessage) Watch(messageId string) interface{} {
	// 获取相应的数据，组织成监控的数据对象，然后返回
	return nil
}

// SpecialUrgencyMessage 特急消息
type SpecialUrgencyMessage struct {
	MessageImplementor
}

func NewSpecialUrgencyMessage(impl MessageImplementor) *SpecialUrgencyMessage {
	return &SpecialUrgencyMessage{
		MessageImplementor: impl,
	}
}

func (m *SpecialUrgencyMessage) hurry(messageId string) {
	// 执行催促的业务，发出催促的消息
}

func (m *SpecialUrgencyMessage) SendMessage(message, toUser string) {
	message += "特急："
	m.MessageImplementor.Send(message, toUser)
	// 还需要增加一条待催促的信息
}