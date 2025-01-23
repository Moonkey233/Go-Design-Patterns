package adapter

import "fmt"

//goland:noinspection GoTestName
func ExampleAdapter() {
	into := 23.3
	MAdapter := NewLengthAdapter().getLength("m", into)
	MLength := NewM().getLength(into)
	fmt.Printf("MAdapter: %v, MLength: %v\n", MAdapter, MLength)

	CmAdapter := NewLengthAdapter().getLength("cm", into)
	CmLength := NewCm().getLength(into)
	fmt.Printf("CmAdapter: %v, CmLength: %v", CmAdapter, CmLength)

	// Output:
	// MAdapter: 2.33, MLength: 2.33
	// CmAdapter: 233, CmLength: 233
}
