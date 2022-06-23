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

	// Print receipt to devices
	GetDevices(receipt)
}

// Get User oder
func GetDrinkOrder() bill.Bill {

	// create new bill
	receipt := bill.NewReceipt()

	for {

		PrintMenu()
		order := GetInput()

		// create new beverage
		beverage := drink.BuildDrink(order)

		// Get topping order
		topping := GetToppingOrder()

		// Mix beverage with topping
		orderDrink := drink.MixBeverageWithTopping(beverage, topping)

		fmt.Println("------Mixing---")
		fmt.Println("Current Cup: ", orderDrink.Description)
		fmt.Println("Current Cup price: ", orderDrink.Cost)

		receipt = bill.AddCupToReceipt(receipt, orderDrink)

		// Singal to stop ordering
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
		// Print current topping and get user order
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
			fmt.Println("Not such a devices")
			continue
		}

		fmt.Println("Current topping: ", topping.GetDescription())

		// Get signal to stop
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
			fmt.Println("Not such a devices")
			continue
		}

		signal := GetSignal()
		if signal == 0 {
			break
		}
	}
	fmt.Println("-----------Final bill----------")
	fmt.Println(receipt.GetDrink())
	fmt.Println("Total cost: ", receipt.GetCost())
	fmt.Println("Print bill to " + receipt.PrintBill())
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
	fmt.Println("Continue (0 for No/ 1 for Yes) ?")
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
}

// Print devices
func PrintDevices() {
	fmt.Println("---------Devices--------")
	fmt.Println("1. Messenger")
	fmt.Println("2. SMS")
	fmt.Println("3. Console")
}
