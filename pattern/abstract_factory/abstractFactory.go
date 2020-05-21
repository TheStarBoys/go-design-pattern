package abstract_factory

import (
	"fmt"
)

// AbstractFactory 抽象工厂接口，声明创建抽象产品对象的操作
type AbstractFactory interface {
	CreateCPUApi() CPUApi // 创建CPU的对象
	CreateMainboardApi() MainboardApi // 创建主板的对象
}

// CPUApi CPU的接口
type CPUApi interface {
	Calculate()
}

// MainboardApi 主板的接口
type MainboardApi interface {
	InstallCPU()
}

// IntelCPU 英特尔CPU实现
type IntelCPU struct {
	Pins int // 针脚数
}

func (c *IntelCPU) Calculate() {
	fmt.Println("now in Intel CPU, pins = ", c.Pins)
}

// AMDCPU AMD CPU的实现
type AMDCPU struct {
	Pins int // 针脚数
}

func (c *AMDCPU) Calculate() {
	fmt.Println("now in AMD CPU, pins = ", c.Pins)
}

// GAMainboard 技嘉的主板
type GAMainboard struct {
	CPUHoles int // CPU插孔数
}

func (m *GAMainboard) InstallCPU() {
	fmt.Println("now in GAMainboard, cpuHoles = ", m.CPUHoles)
}

// MSIMainboard 微星的主板
type MSIMainboard struct {
	CPUHoles int // CPU插孔数
}

func (m *MSIMainboard) InstallCPU() {
	fmt.Println("now in MSIMainboard, cpuHoles = ", m.CPUHoles)
}

// 抽象工厂的实现对象，也就是具体的装机方案

// 装机方案一：Intel的CPU + 技嘉的主板
type Schema1 struct {}

func (s *Schema1) CreateCPUApi() CPUApi {
	return &IntelCPU{Pins:1156}
}

func (s *Schema1) CreateMainboardApi() MainboardApi {
	return &GAMainboard{CPUHoles: 1156}
}

// 装机方案二：AMD的CPU + 微星的主板
type Schema2 struct {}

func (s *Schema2) CreateCPUApi() CPUApi {
	return &AMDCPU{Pins:939}
}

func (s *Schema2) CreateMainboardApi() MainboardApi {
	return &MSIMainboard{CPUHoles:939}
}

// ComputerEngineer 装机工程师
type ComputerEngineer struct {
	cpu 	  CPUApi
	mainboard MainboardApi
}

func (c *ComputerEngineer) MakeComputer(schema AbstractFactory) {
	// 1. 首先准备好装机所需要的配件
	c.prepareHardwares(schema)
	// 2. 组装电脑
	// 3. 测试电脑
	// 4. 交付客户
}

func (c *ComputerEngineer) prepareHardwares(schema AbstractFactory) {
	// 这里要去准备CPU和主板的具体实现，为了示例简单，这里只准备这两个
	// 可是，装机工程师并不知道如何去创建，怎么办呢?

	// 使用抽象工厂来获取相应的接口对象
	c.cpu = schema.CreateCPUApi()
	c.mainboard = schema.CreateMainboardApi()

	// 测试配件是否好用
	c.cpu.Calculate()
	c.mainboard.InstallCPU()
}