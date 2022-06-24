package bill

type Messenger struct {
	Description string
	Cost        float32
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
	return mess.Device.PrintBill() + " Messenger "
}

func (mess *Messenger) GetDrink() string {
	return mess.Device.GetDrink() + mess.Description
}

func (mess *Messenger) GetCost() float32 {
	return mess.Device.GetCost() + mess.Cost
}
