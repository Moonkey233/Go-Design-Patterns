// 模板方法模式

package templatemethod

import "fmt"

type PrintTemplate interface {
	Print(name string)
}

type template struct {
	isTemplate PrintTemplate
	name       string
}

func (t *template) Print() {
	t.isTemplate.Print(t.name)
}

type A struct{}

func (a *A) Print(name string) {
	fmt.Println("a: " + name)
	// todo
}

type B struct{}

func (b *B) Print(name string) {
	fmt.Println("b: " + name)
	// todo
}
