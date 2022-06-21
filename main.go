package main

import (
	"CoffeeBill/drink"
	"fmt"
	"reflect"
)

func main() {
	// Get user choice
	GetOrder()

}

// Get User oder
func GetOrder() {
	// newCup := &DrinkCup{}
	var totalPrice float32 = 0

	for {
		newCup := &drink.DrinkCup{}
		// var totalPrice float32

		PrintMenu()
		order := GetInput()

		switch order {

		// Dark Roast
		case 1:
			cup := &drink.DarkRoast{
				Cup: newCup,
			}
			fmt.Println(cup.GetPrice())

		// Milk
		case 2:
			cup := &drink.Milk{
				Cup: newCup,
			}
			fmt.Println(cup.GetPrice())

		case 3:
			cup := &drink.IceCream{
				Cup: newCup,
			}
			fmt.Println(cup.GetPrice())

		// Not found
		default:
			fmt.Println("Not found drink, please find something else")
			continue
		}

		// Get drink order and build drink by calling director
		drinkType := drink.GetDrink(order)
		director := drink.NewDirector(drinkType)
		drink := director.BuildBeverager()
		totalPrice += drink.Cost

		fmt.Println("Drink cost: ", drink.Cost)
		fmt.Println("Current price: ", totalPrice)

		var signal int
		fmt.Println("Continue ?")
		fmt.Scanf("%d", &signal)

		if signal == 0 {
			break
		}
	}
	fmt.Println("Total price: ", totalPrice)
}

// Print today's menu
func PrintMenu() {
	fmt.Println("-----Today menu-----")
	fmt.Println("1. Dark Roast")
	fmt.Println("2. Milk")
	fmt.Println("3. Ice cream")
}

// Get user input
func GetInput() int {
	var drinkType int
	fmt.Print("What do u want to drink? ")
	fmt.Scanf("%d", &drinkType)
	return drinkType
}

func PrintTopping() {

}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
