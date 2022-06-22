package main

import (
	"CoffeeBill/drink"
	"CoffeeBill/toppings"
	"fmt"
)

func main() {
	// Get user choice
	GetDrinkOrder()

}

// Get User oder
func GetDrinkOrder() {

	// create new topping

	// create new cup
	cup := drink.NewCup()

	for {

		PrintMenu()
		order := GetInput()

		// New topping

		switch order {

		case 1:
			// Fill cup with DarkRoast
			cup = drink.NewDarkRoastWithCup(cup)

		case 2:
			// Fill cup with milk
			cup = drink.NewMilkWithCup(cup)

		case 3:
			// Fill cup with ice cream
			cup = drink.NewIceCreamWithCup(cup)

		default:
			// Fill cup with nothing
			fmt.Print("Not found, choose somethings to drink\n")
			continue
		}

		GetToppingOrder()

		fmt.Println("Current Bill: ", cup.GetDescription())
		fmt.Println("Current price: ", cup.GetPrice())

		signal := GetSignal()
		if signal == 0 {
			break
		}

	}
	fmt.Println("Total price: ", cup.GetPrice())
}

// Get Topping order
func GetToppingOrder() {
	topping := toppings.NewTopping()

	for {
		PrintTopping()
		order := GetInput()

		switch order {
		case 1:
			// Fill topping with cream
			topping = toppings.NewCream(topping)

		case 2:
			// Fill topping with milk foam
			topping = toppings.NewMilkFoam(topping)
		case 3:

		default:
			break
		}

		fmt.Println("Current Bill: ", topping.GetDescription())
		fmt.Println("Current price: ", topping.GetPrice())

		signal := GetSignal()
		if signal == 0 {
			break
		}
	}
}

// Get drink order and build drink by calling director
// Choose drink to build
// After build drink return object drink{description, costs}
func BuildDrink(order int) drink.Drink {
	drinkType := drink.GetDrink(order)
	director := drink.NewDirector(drinkType)
	dR := director.BuildBeverager()
	return dR
}

// Get user input
func GetInput() int {
	var drinkType int
	fmt.Print("What do u want to order? ")
	fmt.Scanf("%d", &drinkType)
	return drinkType
}

// Get signal to stop
func GetSignal() int {
	var signal int
	fmt.Println("Continue ?")
	fmt.Scanf("%d", &signal)
	return signal
}

// Print today's menu
func PrintMenu() {
	fmt.Println("-----Today menu-----")
	fmt.Println("1. Dark Roast")
	fmt.Println("2. Milk")
	fmt.Println("3. Ice cream")
}

// Print topping menu
func PrintTopping() {
	fmt.Println("-----Today Topping-----")
	fmt.Println("1. Cream")
	fmt.Println("2. Milk Foam")
	fmt.Println("3. Cheese")
}
