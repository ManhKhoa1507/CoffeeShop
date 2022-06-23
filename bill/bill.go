package bill

import (
	"CoffeeBill/drink"
	"strconv"
)

type Bill interface {
	PrintBill() string
	GetDrink() string
	GetCost() float32
}

type Receipt struct {
	Drink drink.Drink
}

func (receipt *Receipt) PrintBill() string {
	return ""
}

func (receipt *Receipt) GetDrink() string {
	return receipt.Drink.Description
}

func (receipt *Receipt) GetCost() float32 {
	return receipt.Drink.Cost
}

// Create new receipt
func NewReceipt() Bill {
	drink := &drink.Drink{
		Description: "",
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
	cost := strconv.FormatFloat(float64(beverage.Cost), 'g', 5, 64)
	currentDescription := receipt.GetDrink() + beverage.Description + " with cost " + cost + "\n"
	currentCost := receipt.GetCost() + beverage.Cost

	currentDrink := &drink.Drink{
		Description: currentDescription,
		Cost:        currentCost,
	}

	return &Receipt{
		Drink: *currentDrink,
	}
}
