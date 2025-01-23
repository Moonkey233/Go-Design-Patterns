package memento

func ExampleText() {
	text := &Text{
		content: "how are you?",
	}

	text.Show()
	progress := text.Save()

	text.Write("I'm fine, thank you, and you?")
	text.Show()

	text.Load(progress)
	text.Show()

	// Output:
	// content: how are you?
	// content: I'm fine, thank you, and you?
	// content: how are you?
}
