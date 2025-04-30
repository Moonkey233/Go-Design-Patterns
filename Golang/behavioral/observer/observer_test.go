package observer

func ExampleObserver() {
	subject := NewSubject()

	boy := NewReader("小明")
	girl := NewReader("小美")

	subject.AddObserver(boy)
	subject.AddObserver(girl)

	subject.UpdateContext("hi~")

	// Output:
	// 小明 get hi~
	// 小美 get hi~
}
