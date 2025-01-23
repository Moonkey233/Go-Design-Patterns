package templatemethod

//goland:noinspection GoTestName
func ExampleTemplateMethod() {
	templateA := &A{}
	t := &template{
		isTemplate: templateA,
		name:       "hi~",
	}
	t.Print()

	templateB := &B{}
	t.isTemplate = templateB
	t.Print()

	// Output:
	// a: hi~
	// b: hi~
}
