package bill

import "CoffeeBill/drink"

type Console struct {
	Drink  drink.Drink
	Device Bill
}

func NewConsole() *Console {
	return &Console{}
}

func NewConsoleWithBill(device Bill) *Console {
	return &Console{
		Device: device,
	}
}

func (console *Console) PrintBill() string {
	return console.Device.PrintBill() + "Console "
}

func (console *Console) PrintDrink() string {
	return console.Device.PrintDrink() + console.Drink.GetDescription()
}
