package bill

type Console struct {
	Description string
	Cost        float32
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
	return console.Device.PrintBill() + "Console "
}

func (console *Console) GetDrink() string {
	return console.Device.GetDrink() + console.Description
}

func (console *Console) GetCost() float32 {
	return console.Device.GetCost() + console.Cost
}
