package singleton

import "sync"

type singleton struct {}

var (
	instance *singleton
	once sync.Once
)

func Instance() *singleton {
	// once.Do() 保证了该匿名函数只被执行一次
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
