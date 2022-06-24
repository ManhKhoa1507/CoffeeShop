package drink

type DarkRoast struct {
	Description string
	Cost        float32
}

// Constructor new dark roast
func NewDarkRoast() *DarkRoast {
	return &DarkRoast{}
}

// Set dark roast cost
func (darkRoast *DarkRoast) SetCost() {
	darkRoast.Cost = 20
}

// Set dark roast description
func (darkRoast *DarkRoast) SetDescription() {
	darkRoast.Description = "Dark roast"
}

// Get dark roast
func (darkRoast *DarkRoast) GetBeverager() Drink {
	return Drink{
		Description: darkRoast.Description,
		Cost:        darkRoast.Cost,
	}
}
