package command

//goland:noinspection GoTestName
func ExampleCommand() {
	t := &tv{}
	onTvCommand := &onCommand{
		device: t,
	}
	offTvCommand := &offCommand{
		device: t,
	}
	onTvButton := &button{
		command: onTvCommand,
	}
	offTvButton := &button{
		command: offTvCommand,
	}
	onTvButton.press()
	offTvButton.press()

	ac := &airConditioner{}
	onAirConditionerCommand := &onCommand{
		device: ac,
	}
	offAirConditionerCommand := &offCommand{
		device: ac,
	}
	onAirConditionerButton := &button{
		command: onAirConditionerCommand,
	}
	offAirConditionerButton := &button{
		command: offAirConditionerCommand,
	}
	onAirConditionerButton.press()
	offAirConditionerButton.press()

	// Output:
	// Turning tv on
	// Turning tv off
	// Turning air conditioner on
	// Turning air conditioner off
}
