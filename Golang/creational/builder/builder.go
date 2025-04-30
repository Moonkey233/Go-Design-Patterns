// 建造者模式

package builder

import "fmt"

// Builder 建造者接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

// Director 管理类
type Director struct {
	builder Builder
}

// NewDirector 构造函数
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct 建造
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

type MyBuilder struct {
	material1 string
	material2 string
	material3 string
}

func (b *MyBuilder) Part1() {
	fmt.Printf("part1: %v\n", b.material1)
}

func (b *MyBuilder) Part2() {
	fmt.Printf("part2: %v\n", b.material2)
}

func (b *MyBuilder) Part3() {
	fmt.Printf("part3: %v\n", b.material3)
}
