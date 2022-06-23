package bill

import (
	"CoffeeBill/drink"
	"strconv"
)

type Bill interface {
	PrintBill() string
	PrintDrink() string
}

type Receipt struct {
	Drink drink.Drink
}

func (receipt *Receipt) PrintBill() string {
	return ""
}

func (receipt *Receipt) PrintDrink() string {
	return receipt.Drink.Description
}

// Create new receipt
func NewReceipt() Bill {
	drink := &drink.Drink{
		Description: "Bill: ",
		Cost:        0,
	}

	return &Receipt{
		Drink: *drink,
	}
}

// Add cup to the receipt
// Get old receipt then add cup to current receipt
func AddCupToReceipt(receipt Bill, beverage drink.Drink) Bill {

	// Convert float32 to strings
	currentCost := strconv.FormatFloat(float64(beverage.Cost), 'g', 5, 64)
	currentDescription := receipt.PrintDrink() + beverage.Description + " with cost " + currentCost

	currentDrink := &drink.Drink{
		Description: currentDescription,
	}

	return &Receipt{
		Drink: *currentDrink,
	}
}
