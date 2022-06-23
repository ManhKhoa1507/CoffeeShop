package bill

type Messenger struct {
	Description string
	Device      Bill
}

func NewMess() *Messenger {
	return &Messenger{}
}

func NewMessWithBill(device Bill) *Messenger {
	return &Messenger{
		Description: device.PrintBill(),
		Device:      device,
	}
}

func (mess *Messenger) PrintBill() string {
	return mess.Device.PrintBill() + " Messenger "
}

func (mess *Messenger) PrintDrink() string {
	return mess.Device.PrintDrink() + mess.Description
}
