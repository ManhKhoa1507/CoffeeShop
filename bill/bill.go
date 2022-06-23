package bill

import "CoffeeBill/drink"

type Bill interface {
	PrintBill() string
}

type Receipt struct {
	Drink drink.Drink
}

func (receipt *Receipt) PrintBill() string {
	return receipt.Drink.Description + " print bill from: "
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
func AddCupToReceipt(beverage drink.Drink) Bill {
	return &Receipt{
		Drink: beverage,
	}
}
