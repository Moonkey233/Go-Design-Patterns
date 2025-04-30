// 观察者模式

package observer

import "fmt"

// Subject 发布者
type Subject struct {
	observers []Observer
	content   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

// AddObserver 添加订阅者
func (s *Subject) AddObserver(o Observer) {
	s.observers = append(s.observers, o)
}

// UpdateContext 改变发布者的状态
func (s *Subject) UpdateContext(content string) {
	s.content = content
	s.notify()
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Do(s)
	}
}

// Observer 通知订阅者接口
type Observer interface {
	Do(*Subject)
}

// Reader 订阅者
type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

func (r *Reader) Do(s *Subject) {
	fmt.Println(r.name + " get " + s.content)
}
