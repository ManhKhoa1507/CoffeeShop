package bill

import "CoffeeBill/drink"

type Messenger struct {
	Description string
	Drink       drink.Drink
	Device      Bill
}

func NewMess() *Messenger {
	return &Messenger{}
}

func NewMessWithBill(device Bill) *Messenger {
	return &Messenger{
		Device: device,
	}
}

func (mess *Messenger) PrintBill() string {
	return mess.Device.PrintBill() + mess.Description + " Messenger "
}
