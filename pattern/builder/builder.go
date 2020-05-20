package builder

// Builder 生成器接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

// Product 产品对象接口
// 在使用生成器模式的时候，大多数情况是不知道最终构建出来的产品是什么样的
// 所以在标准的生成器模式里面，一般是不需要对产品定义接口的，因为最终
// 构造的产品千差万别，给这些产品定义公共接口几乎没有意义
//type Product interface {
//	Operate1()
//	Operate2()
//}

// Director 指导者
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder:builder,
	}
}

// Construct 指导者指导生成产品对象
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

// 生成器实现
type Builder1 struct {
	result string
}

func (b *Builder1) Part1() {
	b.result += "1"
}

func (b *Builder1) Part2() {
	b.result += "2"
}

func (b *Builder1) Part3() {
	b.result += "3"
}

func (b *Builder1) GetResult() string {
	return b.result
}


type Builder2 struct {
	result int
}

func (b *Builder2) Part1() {
	b.result += 1
}

func (b *Builder2) Part2() {
	b.result += 2
}

func (b *Builder2) Part3() {
	b.result += 3
}

func (b *Builder2) GetResult() int {
	return b.result
}