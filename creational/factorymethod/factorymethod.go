// 工厂方法模式

package factorymethod

import "fmt"

type Pay interface {
	Pay() string
}

type PayReq struct {
	OrderId string
}

type APayReq struct {
	PayReq
}

func (p *APayReq) Pay() string {
	// todo
	p.OrderId = "ATestId"
	fmt.Println(p.OrderId)
	return "APay支付成功"
}

type BPayReq struct {
	PayReq
	Uid int64
}

func (p *BPayReq) Pay() string {
	// todo
	p.OrderId = "BTestId"
	p.Uid = 123
	fmt.Println(p.OrderId)
	fmt.Println(p.Uid)
	return "BPay支付成功"
}
