package flyweight

//goland:noinspection GoTestName
func ExampleFlyweight() {
	viewer1 := NewColorViewer("blue")
	viewer2 := NewColorViewer("blue")

	PrintIsSameAddress(viewer1, viewer2, "viewer1 address", "viewer2 address")
	PrintIsSameAddress(viewer1.ColorFlyweight, viewer2.ColorFlyweight, "viewer1.ColorFlyweight address", "viewer2.ColorFlyweight address")

	// Output:
	// viewer1 address == viewer2 address: false
	// viewer1.ColorFlyweight address == viewer2.ColorFlyweight address: true
}
