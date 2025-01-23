package strategy

func ExampleTravel() {
	walk := &Walk{}
	Travel1 := NewTravel("小明", walk)
	Travel1.traffic()

	ride := &Ride{}
	Travel2 := NewTravel("小美", ride)
	Travel2.traffic()

	drive := &Drive{}
	Travel3 := NewTravel("小刚", drive)
	Travel3.traffic()

	// Output:
	// 小明 walk
	// 小美 ride
	// 小刚 drive
}
