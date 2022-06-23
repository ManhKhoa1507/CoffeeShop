package bill

type Sms struct {
	Description string
	Cost        float32
	Device      Bill
}

func NewSms() *Sms {
	return &Sms{}
}

func NewSmsWithBill(device Bill) *Sms {
	return &Sms{
		Device: device,
	}
}

func (sms *Sms) PrintBill() string {
	return sms.Device.PrintBill() + "SMS "
}

func (sms *Sms) GetDrink() string {
	return sms.Device.GetDrink() + sms.Description
}

func (sms *Sms) GetCost() float32 {
	return sms.Device.GetCost() + sms.Cost
}
