package mediator

//goland:noinspection GoTestName
func ExampleMediator() {
	message := &Message{}

	p1 := &p1{}
	p2 := &p2{}
	p3 := &p3{}

	message.sendMessage(p1, "hi! my name is p1")
	message.sendMessage(p2, "hi! my name is p2")
	message.sendMessage(p3, "hi! my name is p3")

	// Output:
	// p2 get message: hi! my name is p1
	// p1 get message: hi! my name is p2
	// p1 get message: hi! my name is p3
	// p2 get message: hi! my name is p3
}
