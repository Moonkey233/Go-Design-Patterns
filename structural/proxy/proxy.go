// 代理模式

package proxy

import "fmt"

type Subject interface {
	Proxy() string
}

// Proxy 代理
type Proxy struct {
	real *RealSubject
}

func (p *Proxy) Proxy() string {
	var res string

	// 在调用真实对象之前，检查缓存，判断权限，等等
	p.real.Pre()

	// 调用真实对象
	p.real.Real()

	// 调用之后的操作，如缓存结果，对结果进行处理，等等
	p.real.After()

	return res
}

// RealSubject 真实对象
type RealSubject struct{}

func (*RealSubject) Real() {
	fmt.Print("real")
}

func (*RealSubject) Pre() {
	fmt.Print("pre:")
}

func (*RealSubject) After() {
	fmt.Print(":after\n")
}
