package factorymethod

import "fmt"

//goland:noinspection GoTestName
func ExampleFactoryMethod() {
	aPay := APayReq{}
	bPay := BPayReq{}
	fmt.Println(aPay.Pay())
	fmt.Println(bPay.Pay())

	// Output:
	// ATestId
	// APay支付成功
	// BTestId
	// 123
	// BPay支付成功
}
