package drink

type Milk struct {
	Description string
	Cost        float32
}

// constructor new milk
func NewMilk() *Milk {
	return &Milk{}
}

// Set milk cost
func (milk *Milk) SetCost() {
	milk.Cost = 10
}

// Set milk description
func (milk *Milk) SetDescription() {
	milk.Description = "Milk"
}

// Return milk drink
func (milk *Milk) GetBeverager() Drink {
	return Drink{
		Description: milk.Description,
		Cost:        milk.Cost,
	}
}
