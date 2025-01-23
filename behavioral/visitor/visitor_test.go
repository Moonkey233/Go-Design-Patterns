package visitor

//goland:noinspection GoTestName
func ExampleVisitor() {
	s := &square{}
	c := &circle{}

	side := &sideCalculator{}

	s.accept(side)
	c.accept(side)

	radius := &radiusCalculator{}
	s.accept(radius)
	c.accept(radius)

	// Output:
	// square side
	// circle side
	// square radius
	// circle radius
}
