package state

func ExampleState() {
	d := &door{id: 1}

	// 开启状态
	op := &opened{}
	d.setState(op)
	d.open()
	d.close()

	// 关闭状态
	cl := &closed{}
	d.setState(cl)
	d.open()
	d.close()

	// 损坏状态
	dm := &damaged{}
	d.setState(dm)
	d.open()
	d.close()

	// Output:
	// 1门已开启
	// 1关闭成功
	// 1开启成功
	// 1门已关闭
	// 1门已损坏，无法开启
	// 1门已损坏，无法关闭
}
