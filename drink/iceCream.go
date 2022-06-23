package drink

type IceCream struct {
	Description string
	Cost        float32
	Cup         Cup
}

// Constructor new ice cream
func NewIceCream() *IceCream {
	return &IceCream{
		Description: "ice cream",
		Cost:        15,
	}
}

// Fill cup with Ice cream
func NewIceCreamWithCup(cup Cup) *IceCream {
	return &IceCream{
		Description: "ice cream",
		Cost:        15,
		Cup:         cup,
	}
}

func (iceCream *IceCream) GetPrice() float32 {
	cupPrize := iceCream.Cup.GetPrice()
	return cupPrize + iceCream.Cost
}

func (iceCream *IceCream) GetDescription() string {
	cupDescription := iceCream.Cup.GetDescription()
	return cupDescription + " " + "ice cream"
}
