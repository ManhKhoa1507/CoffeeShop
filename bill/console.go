package bill

import "CoffeeBill/drink"

type Console struct {
	Description string
	Drink       drink.Drink
	Device      Bill
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
	return console.Device.PrintBill() + console.Drink.GetDescription() + "Console "
}
