// 适配器模式

package adapter

import "fmt"

// Cm 提供一个获取厘米的接口
type Cm interface {
	getLength(float64) float64
}

// M 提供一个获取米的接口
type M interface {
	getLength(float64) float64
}

func NewM() M {
	return &LengthM{}
}

type LengthM struct{}

func (*LengthM) getLength(cm float64) float64 {
	return cm / 10
}

func NewCm() Cm {
	return &LengthCm{}
}

type LengthCm struct{}

func (a *LengthCm) getLength(m float64) float64 {
	return m * 10
}

// LengthAdapter 适配器
type LengthAdapter interface {
	getLength(string, float64) float64
}

func NewLengthAdapter() LengthAdapter {
	return &ExtendLengthAdapter{}
}

type ExtendLengthAdapter struct{}

func (*ExtendLengthAdapter) getLength(isType string, into float64) float64 {
	if isType == "m" {
		return NewM().getLength(into)
	} else if isType == "cm" {
		return NewCm().getLength(into)
	}
	fmt.Printf("unknown type: %v", isType)
	return into
}
