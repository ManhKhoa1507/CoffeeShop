package toppings

type Cream struct {
	Description string
	Cost        float32
	Topping     Topping
}

// Constructor new cream
func NewCream(topping Topping) *Cream {
	return &Cream{
		Description: "cream",
		Cost:        1,
		Topping:     topping,
	}
}

// Get Cream function
// Get Cream price

func (cream *Cream) GetPrice() float32 {
	topPrice := cream.Topping.GetPrice()
	return topPrice + 1
}

func (cream *Cream) GetDescription() string {
	topDescription := cream.Topping.GetDescription()
	return topDescription + " " + cream.Description
}
