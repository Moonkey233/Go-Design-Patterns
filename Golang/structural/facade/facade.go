// 外观模式

package facade

import "fmt"

// APIA 初始化APIA
type APIA interface {
	TestA() string
}

func NewAPIA() APIA {
	return &apiRunA{}
}

type apiRunA struct{}

func (*apiRunA) TestA() string {
	return "A api running"
}

// APIB 初始化APIB
type APIB interface {
	TestB() string
}

func NewAPIB() APIB {
	return &apiRunB{}
}

type apiRunB struct{}

func (*apiRunB) TestB() string {
	return "B api running"
}

// API 外观类
type API interface {
	Test() string
}

func NewAPI() API {
	return &apiRun{
		a: NewAPIA(),
		b: NewAPIB(),
	}
}

type apiRun struct {
	a APIA
	b APIB
}

func (a *apiRun) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}
