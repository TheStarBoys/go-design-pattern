package mediator

import "testing"

func TestMediator(t *testing.T) {
	// 1. 创建中介者——主板对象
	mediator := new(MotherBoard)
	// 2. 创建同事类
	cd := NewCDDriver(mediator)
	cpu := NewCPU(mediator)
	vc := NewVideoCard(mediator)
	sc := NewSoundCard(mediator)

	// 3. 让中介者知道所有同事
	mediator.CDDriver = cd
	mediator.CPU = cpu
	mediator.VideoCard = vc
	mediator.SoundCard = sc

	// 4. 开始看电影，把光盘放入光驱，光驱开始读盘
	cd.ReadCD()
}
