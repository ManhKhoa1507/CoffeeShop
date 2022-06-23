package main

import (
	"CoffeeBill/bill"
	"CoffeeBill/drink"
	"CoffeeBill/toppings"
	"fmt"
)

func main() {
	// Get user choice
	receipt := GetDrinkOrder()
	GetDevices(receipt)

	// Add receipt to devices
}

// Get User oder
func GetDrinkOrder() bill.Bill {

	// create new bill
	receipt := bill.NewReceipt()

	for {

		PrintMenu()
		order := GetInput()

		// create new beverage
		beverage := drink.NewCup()

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

		// Get topping order
		topping := GetToppingOrder()

		// Mix beverage with topping
		orderDrink := drink.MixBeverageWithTopping(beverage, topping)

		fmt.Println("------Mixing---")
		fmt.Println("Current Cup: ", orderDrink.Description)
		fmt.Println("Current Cup price: ", orderDrink.Cost)

		receipt = bill.AddCupToReceipt(orderDrink)

		signal := GetSignal()
		if signal == 0 {
			break
		}

	}

	// Return drink
	return receipt
}

// Get Topping order
func GetToppingOrder() toppings.Topping {

	// Create new topping
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

func GetDevices(receipt bill.Bill) {
	for {
		// Choose devices to print
		PrintDevices()
		order := GetInput()

		switch order {
		case 1:
			receipt = bill.NewMessWithBill(receipt)
		case 2:
			receipt = bill.NewSmsWithBill(receipt)
		case 3:
			receipt = bill.NewConsoleWithBill(receipt)
		default:
			break
		}

		signal := GetSignal()
		if signal == 0 {
			break
		}
	}
	fmt.Println(receipt.PrintBill())
}

// Get user input
func GetInput() int {
	var order int
	fmt.Print("What do u want to order? ")
	fmt.Scanf("%d", &order)
	return order
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

// Print devices
func PrintDevices() {
	fmt.Println("---------Devices--------")
	fmt.Println("1. Messenger")
	fmt.Println("2. SMS")
	fmt.Println("3. Console")
}
