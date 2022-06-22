package drink

type DarkRoast struct {
	Description string
	Cost        float32
	Cup         Cup
}

// Constructor new dark roast
func NewDarkRoast() *DarkRoast {
	return &DarkRoast{}
}

func NewDarkRoastWithCup(cup Cup) *DarkRoast {
	return &DarkRoast{
		Cup: cup,
	}
}

// Set dark roast cost
func (darkRoast *DarkRoast) SetCost() {
	darkRoast.Cost = 20
}

// Set dark roast description
func (darkRoast *DarkRoast) SetDescription() {
	darkRoast.Description = "Dark roast"
}

func (darkRoast *DarkRoast) GetPrice() float32 {
	cupPrice := darkRoast.Cup.GetPrice()
	return cupPrice + 20
}

// Get darkRoast description
func (darkRoast *DarkRoast) GetDescription() string {
	cupDescription := darkRoast.Cup.GetDescription()
	return cupDescription + " " + "dark roast"
}

// Get dark roast
func (darkRoast *DarkRoast) GetBeverager() Drink {
	return Drink{
		Description: darkRoast.Description,
		Cost:        darkRoast.Cost,
	}
}
