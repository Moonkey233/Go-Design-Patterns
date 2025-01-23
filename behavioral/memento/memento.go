// 备忘录模式

package memento

import "fmt"

type Memento interface{}

type Text struct {
	content string
}

type textMemento struct {
	content string
}

func (t *Text) Write(content string) {
	t.content = content
}

func (t *Text) Save() Memento {
	return &textMemento{
		content: t.content,
	}
}

func (t *Text) Load(m Memento) {
	tm := m.(*textMemento)
	t.content = tm.content
}

func (t *Text) Show() {
	fmt.Println("content:", t.content)
}
