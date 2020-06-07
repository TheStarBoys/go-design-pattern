package flyweight

// Flyweight 享元接口，通过这个接口享元可以接受并作用于外部状态
type Flyweight interface {
	// 传入外部状态
	Operation(extrinsicState string)
}

// ConcreteFlyweight 享元对象
type ConcreteFlyweight struct {
	intrinsicState string // 内部状态
}

func (f *ConcreteFlyweight) Operation(extrinsicState string) {
	// 具体功能处理
}

// UnsharedConcreteFlyweight 不需要共享的 flyweight对象
// 通常是被共享的享元对象作为子节点组合出来的对象
type UnsharedConcreteFlyweight struct {
	allState string
}

func (f *UnsharedConcreteFlyweight) Operation(extrinsicState string) {
	// 具体功能处理
}

// 享元工厂
type FlyweightFactory struct {
	// 缓存多个 flyweight对象
	fsMap map[string]Flyweight
}

func (factory *FlyweightFactory) GetFlyweight(key string) Flyweight {
	f, ok := factory.fsMap[key]
	if !ok {
		f = &ConcreteFlyweight{
			intrinsicState: "some state",
		}
		factory.fsMap[key] = f
	}

	return f
}