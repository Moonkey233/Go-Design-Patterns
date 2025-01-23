package iterator

import "fmt"

//goland:noinspection GoTestName
func ExampleIterator() {
	var pc collection
	part1 := &part{
		title:  "part1",
		number: 10,
	}
	part2 := &part{
		title:  "part2",
		number: 20,
	}
	part3 := &part{
		title:  "part3",
		number: 30,
	}
	pc = &partCollection{
		parts: []*part{part1, part2, part3},
	}
	it := pc.CreateIterator()
	reIt := pc.CreateReIterator()

	for it.HasNext() {
		p := it.GetNext()
		fmt.Println(p)
	}
	for reIt.HasNext() {
		p := reIt.GetNext()
		fmt.Println(p)
	}

	// Output:
	// &{part1 10}
	// &{part2 20}
	// &{part3 30}
	// &{part3 30}
	// &{part2 20}
	// &{part1 10}
}
