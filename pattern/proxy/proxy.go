package proxy

import "fmt"

// Subject 抽象的目标接口，定义具体的目标对象和代理公用的接口
type Subject interface {
	// 一个抽象的请求方法
	Request()
}

// RealSubject 具体的目标对象，是真正被代理的对象
type RealSubject struct {}

func (s *RealSubject) Request() {
	// 执行具体的功能处理
	fmt.Println("request")
}

// Proxy 代理对象
type Proxy struct {
	Real *RealSubject // 持有被代理的具体的目标对象
}

func (p *Proxy) Request() {
	// 在转调具体的目标对象前，可以执行一些功能处理
	fmt.Println("before")
	// 转调具体的目标对象的方法
	p.Real.Request()
	// 在转调具体的目标对象后，可以执行一些功能处理
	fmt.Println("after")
}