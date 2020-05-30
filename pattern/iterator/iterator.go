package iterator

// Iterator 迭代器接口，定义访问和遍历元素的操作
type Iterator interface {
	// 移动到聚合对象的第一个位置
	First()
	// 移动到聚合对象的下一个位置
	Next()
	// 判断是否已经移动到聚合对象的最后一个位置
	IsEnd() bool
	// 获取迭代的当前元素
	CurrentItem() interface{}
}

// ConcreteIterator 具体的迭代器实现对象
type ConcreteIterator struct {
	// 持有被迭代的具体的聚合对象
	aggregate ConcreteAggregate
	// 内部索引，记录当前迭代到的索引位置
	// -1 表示刚开始的时候，迭代器指向聚合对象第一个对象之前
	index int
}

func NewConcreteIterator(aggregate ConcreteAggregate) *ConcreteIterator {
	return &ConcreteIterator{
		aggregate: aggregate,
		index: -1,
	}
}

func (c *ConcreteIterator) First() {
	c.index = 0
}

func (c *ConcreteIterator) Next() {
	if c.index < c.aggregate.Size() {
		c.index++
	}
}

func (c *ConcreteIterator) IsEnd() bool {
	if c.index == c.aggregate.Size() {
		return true
	}

	return false
}

func (c *ConcreteIterator) CurrentItem() interface{} {
	return c.aggregate.Get(c.index)
}

// Aggregate 聚合对象的接口，定义创建相应迭代器对象的接口
type Aggregate interface {
	// 工厂方法，创建相应迭代器对象的接口
	CreateIterator() Iterator
}

// ConcreteAggregate 具体的聚合对象，实现创建相应迭代器对象的功能
type ConcreteAggregate struct {
	// 聚合对象具体的内容
	data []string
}

func NewConcreteAggregate(data []string) ConcreteAggregate {
	return ConcreteAggregate{
		data: data,
	}
}

// CreateIterator 实现创建Iterator的工厂方法
func (c ConcreteAggregate) CreateIterator() Iterator {
	return NewConcreteIterator(c)
}

// Get 获取索引对应的元素
func (c ConcreteAggregate) Get(index int) interface{} {
	if index < len(c.data) {
		return c.data[index]
	}

	return nil
}

// Size 获取聚合对象的大小
func (c ConcreteAggregate) Size() int {
	return len(c.data)
}