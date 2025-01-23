// 策略模式

package strategy

import "fmt"

type Travel struct {
	name     string
	strategy Strategy
}

func NewTravel(name string, strategy Strategy) *Travel {
	return &Travel{
		name:     name,
		strategy: strategy,
	}
}

func (p *Travel) traffic() {
	p.strategy.traffic(p)
}

type Strategy interface {
	traffic(*Travel)
}

type Walk struct{}

func (w *Walk) traffic(t *Travel) {
	fmt.Println(t.name + " walk")
}

type Ride struct{}

func (w *Ride) traffic(t *Travel) {
	fmt.Println(t.name + " ride")
}

type Drive struct{}

func (w *Drive) traffic(t *Travel) {
	fmt.Println(t.name + " drive")
}
