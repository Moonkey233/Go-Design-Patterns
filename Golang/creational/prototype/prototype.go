// 原型模式

package prototype

type Type1 struct {
	name string
}

func (t *Type1) Clone() *Type1 {
	tc := *t
	return &tc
}
