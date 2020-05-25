package mediator

import (
	"strings"
	"fmt"
)

// Mediator 中介者，定义各个同事对象通信的接口
type Mediator interface {
	// 同事对象在自身改变的时候来通知中介者的方法
	// 让中介者去负责相应的与其他同事对象的交互
	// colleague 同事对象自身，好让中介者对象通过对象实例去获取同事对象的状态
	Changed(colleague ColleagueInterface)
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

// CDDriver 光驱类，一个同时类
type CDDriver struct {
	Colleague
	data string // 从光驱读出来的数据
}

func NewCDDriver(mediator Mediator) *CDDriver {
	return &CDDriver{
		Colleague: Colleague{
			mediator: mediator,
		},
	}
}

// 获取光驱读取出来的数据
func (cd *CDDriver) GetData() string {
	return cd.data
}

// 读取光盘，业务方法，也是和中介者交互的方法
func (cd *CDDriver) ReadCD() {
	// 逗号前是视频显示的数据，逗号后是声音
	cd.data = "设计模式,值得好好研究"
	// 通知主板，自己的状态发生了改变
	cd.GetMediator().Changed(cd)
}

// CPU类，一个同事类
type CPU struct {
	Colleague
	videoData string // 分解出来的视频数据
	soundData string // 分解出来的声音数据
}

func NewCPU(mediator Mediator) *CPU {
	return &CPU{
		Colleague: Colleague{
			mediator: mediator,
		},
	}
}

// 获取视频数据
func (c *CPU) GetVideoData() string {
	return c.videoData
}

// 获取音频数据
func (c *CPU) GetSoundData() string {
	return c.soundData
}

// 处理数据，把数据分成音频和视频数据
func (c *CPU) ExecuteData(data string) {
	// 把数据分解开，前面的是视频数据，后面的是音频数据
	ss := strings.Split(data, ",")
	c.videoData = ss[0]
	c.soundData = ss[1]
	// 通知主板，CPU的工作完成
	c.GetMediator().Changed(c)
}

// VideoCard 显卡类，一个同事类
type VideoCard struct {
	Colleague
}

func NewVideoCard(mediator Mediator) *VideoCard {
	return &VideoCard{
		Colleague: Colleague{
			mediator: mediator,
		},
	}
}

// 显示视频数据
func (v *VideoCard) ShowData(data string) {
	fmt.Println("您正观看的是：", data)
}

// SoundCard 声卡类，一个同事类
type SoundCard struct {
	Colleague
}

func NewSoundCard(mediator Mediator) *SoundCard {
	return &SoundCard{
		Colleague: Colleague{
			mediator: mediator,
		},
	}
}

// 按照音频数据发出声音
func (s *SoundCard) SoundData(data string) {
	fmt.Println("画外音：", data)
}

// MotherBoard 主板类，实现中介者接口
type MotherBoard struct {
	CDDriver  *CDDriver
	CPU       *CPU
	VideoCard *VideoCard
	SoundCard *SoundCard
}

func (m *MotherBoard) Changed(colleague ColleagueInterface) {
	switch colleague.(type) {
	case *CDDriver:
		// 表示光驱读取数据了
		m.OpeCDDriverReadData()
	case *CPU:
		// 表示CPU处理完了
		m.OpeCPU()
	}
}

// 处理光驱读取数据以后，与其他对象的交互
func (m *MotherBoard) OpeCDDriverReadData() {
	// 1. 先获取光驱读取的数据
	data := m.CDDriver.GetData()
	// 2. 把这些数据传递给CPU进行处理
	m.CPU.ExecuteData(data)
}

// 处理CPU处理完数据后，与其他对象的交互
func (m *MotherBoard) OpeCPU() {
	// 1. 先获取CPU处理后的数据
	videoData := m.CPU.GetVideoData()
	soundData := m.CPU.GetSoundData()
	// 2. 把这些数据传递给显卡和声卡展示出来
	m.VideoCard.ShowData(videoData)
	m.SoundCard.SoundData(soundData)
}