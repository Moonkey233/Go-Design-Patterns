// 访问者模式

package visitor

import "fmt"

type Shape interface {
	getName() string
	accept(visitor)
}

type square struct{}

func (s *square) getName() string {
	return "square"
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

type circle struct{}

func (c *circle) getName() string {
	return "circle"
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
}

type sideCalculator struct{}

func (a *sideCalculator) visitForSquare(s *square) {
	fmt.Printf("%v side\n", s.getName())
}

func (a *sideCalculator) visitForCircle(c *circle) {
	fmt.Printf("%v side\n", c.getName())
}

type radiusCalculator struct{}

func (a *radiusCalculator) visitForSquare(s *square) {
	fmt.Printf("%v radius\n", s.getName())
}

func (a *radiusCalculator) visitForCircle(c *circle) {
	fmt.Printf("%v radius\n", c.getName())
}
