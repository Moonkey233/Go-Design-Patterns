// 桥接模式

package bridge

import "fmt"

// SendMessage 两种发送消息的方法
type SendMessage interface {
	baseSend(text, to string)
}

type sms struct{}

func NewSms() SendMessage {
	return &sms{}
}

func (*sms) baseSend(text, to string) {
	fmt.Println(fmt.Sprintf("send %s to %s by sms", text, to))
}

type email struct{}

func NewEmail() SendMessage {
	return &email{}
}

func (*email) baseSend(text, to string) {
	fmt.Println(fmt.Sprintf("send %s to %s by email", text, to))
}

type SystemA struct {
	method SendMessage
}

func NewSystemA(method SendMessage) *SystemA {
	return &SystemA{
		method: method,
	}
}

func (m *SystemA) SendMessage(text, to string) {
	m.method.baseSend(fmt.Sprintf("[System A] %s", text), to)
}

type SystemB struct {
	method SendMessage
}

func NewSystemB(method SendMessage) *SystemB {
	return &SystemB{
		method: method,
	}
}

func (m *SystemB) SendMessage(text, to string) {
	m.method.baseSend(fmt.Sprintf("[System B] %s", text), to)
}
