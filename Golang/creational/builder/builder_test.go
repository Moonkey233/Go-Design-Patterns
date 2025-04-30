package builder

func ExampleBuilder() {
	builder := &MyBuilder{
		material1: "material1",
		material2: "material2",
		material3: "material3",
	}
	director := NewDirector(builder)
	director.Construct()

	// Output:
	// part1: material1
	// part2: material2
	// part3: material3
}
