package state

func ExampleState() {
	d := &door{id: 1}
	d.setState(&opened{})

	// 开启状态
	d.open()
	d.close()

	// 关闭状态
	d.close()
	d.open()

	// 损坏状态
	d.setState(&damaged{})
	d.open()
	d.close()

	// Output:
	// 1门已开启
	// 1关闭成功
	// 1门已关闭
	// 1开启成功
	// 1门已损坏，无法开启
	// 1门已损坏，无法关闭
}
