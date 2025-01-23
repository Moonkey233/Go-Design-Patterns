package singleton

import "fmt"

func ExampleSingleton() {
	ins1 := GetInstance(1)
	fmt.Println(ins1)
	fmt.Println(ins1.getValue())
	ins2 := GetInstance(2)
	fmt.Println(ins2)
	fmt.Println(ins2.getValue())

	// Output:
	// &{1}
	// 1
	// &{1}
	// 1
}
