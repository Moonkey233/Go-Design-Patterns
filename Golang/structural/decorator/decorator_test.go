package decorator

import (
	"fmt"
)

//goland:noinspection GoTestName
func ExampleDecorator() {

	basePizza := &base{}
	fmt.Printf("basePizza price is %d\n", basePizza.getPrice())

	//Only add tomato topping
	pizzaWithTomato := &tomatoTopping{
		pizza: basePizza,
	}
	fmt.Printf("pizzaWithTomato price is %d\n", pizzaWithTomato.getPrice())

	//Only add cheese topping
	pizzaWithCheese := &cheeseTopping{
		pizza: basePizza,
	}
	fmt.Printf("pizzaWithCheese price is %d\n", pizzaWithCheese.getPrice())

	//Add cheese topping and tomato topping
	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithCheese,
	}
	fmt.Printf("pizzaWithCheeseAndTomato price is %d\n", pizzaWithCheeseAndTomato.getPrice())

	// Output:
	// basePizza price is 15
	// pizzaWithTomato price is 25
	// pizzaWithCheese price is 35
	// pizzaWithCheeseAndTomato price is 45
}
