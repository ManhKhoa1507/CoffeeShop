package drink

type Milk struct {
	Description string
	Cost        float32
	Cup         Cup
}

// constructor new milk
func NewMilk() *Milk {
	return &Milk{}
}

func NewMilkWithCup(cup Cup) *Milk {
	return &Milk{
		Cup: cup,
	}
}

// Set milk cost
func (milk *Milk) SetCost() {
	milk.Cost = 10
}

// Set milk description
func (milk *Milk) SetDescription() {
	milk.Description = "milk"
}

func (milk *Milk) GetPrice() float32 {
	cupPrice := milk.Cup.GetPrice()
	return cupPrice + 10
}

func (milk *Milk) GetDescription() string {
	cupDescription := milk.Cup.GetDescription()
	return cupDescription + " " + "milk"
}

// Return milk drink
func (milk *Milk) GetBeverager() Drink {
	return Drink{
		Description: milk.Description,
		Cost:        milk.Cost,
	}
}
