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

	// create new beverage
	beverage := drink.NewCup()

	for {

		PrintMenu()
		order := GetInput()

		// New topping

		switch order {

		case 1:
			// Fill cup with DarkRoast
			beverage = drink.NewDarkRoastWithCup(beverage)

		case 2:
			// Fill cup with milk
			beverage = drink.NewMilkWithCup(beverage)

		case 3:
			// Fill cup with ice cream
			beverage = drink.NewIceCreamWithCup(beverage)

		default:
			// Fill cup with nothing
			fmt.Print("Not found, choose somethings to drink\n")
			continue
		}

		topping := GetToppingOrder()
		cup := drink.MixBeverageWithTopping(beverage, topping)

		fmt.Println("------Mixing---")
		fmt.Println("Current Bill: ", cup.Description)
		fmt.Println("Current price: ", cup.Cost)

		signal := GetSignal()
		if signal == 0 {
			break
		}

	}
}

// Get Topping order
func GetToppingOrder() toppings.Topping {
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

		fmt.Println("Current topping: ", topping.GetDescription())

		signal := GetSignal()
		if signal == 0 {
			break
		}
	}
	return topping
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
