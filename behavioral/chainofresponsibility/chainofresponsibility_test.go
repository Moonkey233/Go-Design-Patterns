package chainofresponsibility

//goland:noinspection GoTestName
func ExampleChainOfResponsibility() {
	var start Department

	a := &aPart{}
	b := &bPart{}
	end := &endPart{}
	start = b
	a.setNext(end)
	b.setNext(a)
	do := &Do{}
	start.execute(do)

	a = &aPart{}
	b = &bPart{}
	start = a
	end = &endPart{}
	a.setNext(b)
	b.setNext(end)
	do = &Do{}
	start.execute(do)

	// Output:
	// bPart
	// aPart
	// endPart
	// aPart
	// bPart
	// endPart
}
