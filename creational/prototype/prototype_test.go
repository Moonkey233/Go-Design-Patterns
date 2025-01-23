package prototype

import (
	"fmt"
)

//goland:noinspection GoTestName
func ExamplePrototype() {
	t1 := &Type1{
		name: "type1",
	}
	t2 := t1.Clone()
	fmt.Printf("t1 == t2: %v\n", t1 == t2)
	fmt.Println(t1.name)
	fmt.Println(t2.name)
	fmt.Println(t1)
	fmt.Println(t2)
	t2.name = "type2"
	fmt.Println(t1.name)
	fmt.Println(t2.name)
	fmt.Println(t1)
	fmt.Println(t2)

	// Output:
	// t1 == t2: false
	// type1
	// type1
	// &{type1}
	// &{type1}
	// type1
	// type2
	// &{type1}
	// &{type2}
}
