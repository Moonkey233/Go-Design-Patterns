// 单例模式

package singleton

import (
	"sync"
)

// 单例实例
type singleton struct {
	Value int
}

type Singleton interface {
	getValue() int
}

func (s singleton) getValue() int {
	return s.Value
}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 构造方法，用于获取单例模式对象
func GetInstance(v int) Singleton {
	once.Do(func() {
		instance = &singleton{Value: v}
	})
	return instance
}
