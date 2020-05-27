package proxy

import "testing"

func TestProxy(t *testing.T) {
	// 生成具体的被代理对象
	sub := new(RealSubject)
	// 创建代理对象
	proxy := &Proxy{
		Real: sub,
	}
	// 通过代理对象去发起请求
	proxy.Request()
	/*

	output:
		before
		request
		after

	 */
}
