package drink

type IceCream struct {
	Description string
	Cost        float32
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

// Get ice cream
func (iceCream *IceCream) GetBeverager() Drink {
	return Drink{
		Description: iceCream.Description,
		Cost:        iceCream.Cost,
	}
}
