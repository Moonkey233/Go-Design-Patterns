package proxy

func ExampleProxy() {
	var sub Subject
	sub = &Proxy{}
	sub.Proxy()

	// Output:
	// pre:real:after
}
