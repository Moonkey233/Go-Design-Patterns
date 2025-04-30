package composite

import (
	"fmt"
)

func ExampleComposite() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic, reason: %v\n", r)
		}
	}()

	root := NewComponent(CompositeNode, "root")
	c1 := NewComponent(CompositeNode, "c1")
	c2 := NewComponent(CompositeNode, "c2")
	c3 := NewComponent(CompositeNode, "c3")

	l1 := NewComponent(LeafNode, "l1")
	l2 := NewComponent(LeafNode, "l2")
	l3 := NewComponent(LeafNode, "l3")

	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(c3)
	c1.AddChild(l1)
	c2.AddChild(l2)
	c2.AddChild(l3)

	root.Search("")
	// 测试panic & recover
	l4 := NewComponent(2, "l4")
	// 不可到达的代码
	println(l4)

	// Output:
	// +root
	//  +c1
	//   +c3
	//   -l1
	//  +c2
	//   -l2
	//   -l3
	// panic, reason: invalid component kind
}
