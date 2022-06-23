package drink

type DarkRoast struct {
	Description string
	Cost        float32
	Cup         Cup
}

// Constructor new dark roast
func NewDarkRoast() *DarkRoast {
	return &DarkRoast{
		Description: "dark roast",
		Cost:        20,
	}
}

// Fill current cup with dark roast
func NewDarkRoastWithCup(cup Cup) *DarkRoast {
	return &DarkRoast{
		Description: "dark roast",
		Cost:        20,
		Cup:         cup,
	}
}

// Get the current cup price then add wuth darkroast coast
func (darkRoast *DarkRoast) GetPrice() float32 {
	return darkRoast.Cup.GetPrice() + darkRoast.Cost
}

// Get darkRoast description
func (darkRoast *DarkRoast) GetDescription() string {
	cupDescription := darkRoast.Cup.GetDescription()
	return cupDescription + " " + darkRoast.Description
}
