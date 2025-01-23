package facade

import "fmt"

//goland:noinspection GoTestName
func ExampleFacade() {
	api := NewAPI()
	fmt.Println(api.Test())

	// Output:
	// A api running
	// B api running
}
