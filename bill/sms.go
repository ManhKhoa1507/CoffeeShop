package bill

import "CoffeeBill/drink"

type Sms struct {
	Drink  drink.Drink
	Device Bill
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

func (sms *Sms) PrintDrink() string {
	return sms.Device.PrintDrink() + sms.Drink.GetDescription()
}
