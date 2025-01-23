package bridge

//goland:noinspection GoTestName
func ExampleBridge() {
	NewSystemA(NewSms()).SendMessage("hi", "baby")
	NewSystemA(NewEmail()).SendMessage("hi", "baby")
	NewSystemB(NewSms()).SendMessage("hi", "baby")
	NewSystemB(NewEmail()).SendMessage("hi", "baby")

	// Output:
	// send [System A] hi to baby by sms
	// send [System A] hi to baby by email
	// send [System B] hi to baby by sms
	// send [System B] hi to baby by email
}
