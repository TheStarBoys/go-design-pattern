package bridge

import "testing"

func TestBridge(t *testing.T) {
	var (
		impl MessageImplementor
		m MessageImplementor
	)
	// 创建具体的实现
	impl = new(MessageSMS)
	// 创建普通消息对象
	m = NewCommonMessage(impl)
	m.Send("请喝一杯茶", "小李")
	// 创建紧急消息对象
	m = NewUrgencyMessage(impl)
	m.Send("请喝一杯茶", "小李")
	// 创建特急消息对象
	m = NewSpecialUrgencyMessage(impl)
	m.Send("请喝一杯茶", "小李")

	// 把实现方式切换成手机
	impl = new(MessageMobile)
	m = NewCommonMessage(impl)
	m.Send("请喝一杯茶", "小李")
	m = NewUrgencyMessage(impl)
	m.Send("请喝一杯茶", "小李")
	m = NewSpecialUrgencyMessage(impl)
	m.Send("请喝一杯茶", "小李")

	// Output:
	// 使用站内短信方式，发送消息'请喝一杯茶'给小李
	// 使用站内短信方式，发送消息'请喝一杯茶'给小李
	// 使用站内短信方式，发送消息'请喝一杯茶'给小李
	// 使用手机短信方式，发送消息'请喝一杯茶'给小李
	// 使用手机短信方式，发送消息'请喝一杯茶'给小李
	// 使用手机短信方式，发送消息'请喝一杯茶'给小李
}