// 状态模式

package state

import "fmt"

// state 不同状态需要实现的接口
type state interface {
	open(*door)
	close(*door)
}

// door 门对象
type door struct {
	id           int
	currentState state // 当前状态
}

func (d *door) open() {
	d.currentState.open(d)
}

func (d *door) close() {
	d.currentState.close(d)
}

func (d *door) setState(s state) {
	d.currentState = s
}

// opened 开启状态
type opened struct{}

func (o *opened) open(d *door) {
	fmt.Printf("%v门已开启\n", d.id)
}

func (o *opened) close(d *door) {
	fmt.Printf("%v关闭成功\n", d.id)
	d.setState(&closed{})
}

// closed 关闭状态
type closed struct{}

func (c *closed) open(d *door) {
	fmt.Printf("%v开启成功\n", d.id)
	d.setState(&opened{})
}

func (c *closed) close(d *door) {
	fmt.Printf("%v门已关闭\n", d.id)
}

// damaged 损坏状态
type damaged struct{}

func (a *damaged) open(d *door) {
	fmt.Printf("%v门已损坏，无法开启\n", d.id)
}

func (a *damaged) close(d *door) {
	fmt.Printf("%v门已损坏，无法关闭\n", d.id)
}
