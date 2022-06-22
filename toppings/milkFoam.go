package toppings

type MilkFoam struct {
	Description string
	Cost        float32
	Topping     Topping
}

// Constructor new milkFoam
func NewMilkFoam(topping Topping) *MilkFoam {
	return &MilkFoam{
		Description: " Milk foam",
		Cost:        2,
		Topping:     topping,
	}
}

// Get MilkFoam function
// Get price after getting milkfoam
func (milkFoam *MilkFoam) GetPrice() float32 {
	topPrice := milkFoam.Topping.GetPrice()
	return topPrice + milkFoam.Cost
}

// Get top description, after getting milkfoam
func (milkFoam *MilkFoam) GetDescription() string {
	topDescription := milkFoam.Topping.GetDescription()
	return topDescription + " " + milkFoam.Description
}

// Get the milk foam topping
// Return Top struct
func (milkFoam *MilkFoam) GetTopping() Top {
	return Top{
		Description: milkFoam.Description,
		Cost:        milkFoam.Cost,
	}
}
