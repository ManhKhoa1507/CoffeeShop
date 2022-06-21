package drink

type IceCream struct {
	Description string
	Cost        float32
	Cup         Cup
}

// Constructor new ice cream
func NewIceCream() *IceCream {
	return &IceCream{}
}

// Set ice cream cost
func (iceCream *IceCream) SetCost() {
	iceCream.Cost = 15
}

// Set ice cream description
func (iceCream *IceCream) SetDescription() {
	iceCream.Description = "Ice cream"
}

func (iceCream *IceCream) GetCost() float32 {
	return iceCream.Cost
}

func (iceCream *IceCream) GetPrice() float32 {
	cupPrize := iceCream.Cup.GetPrice()
	return cupPrize + iceCream.Cost
}

// Get ice cream
func (iceCream *IceCream) GetBeverager() Drink {
	return Drink{
		Description: iceCream.Description,
		Cost:        iceCream.Cost,
	}
}
