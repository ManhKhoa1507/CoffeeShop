package drink

import (
	"CoffeeBill/toppings"
)

type Cup interface {
	GetDescription() string
	GetPrice() float32
}

func (drink *Drink) GetPrice() float32 {
	return 0
}

func (drink *Drink) GetDescription() string {
	return ""
}

func NewCup() Cup {
	return &Drink{
		Description: "",
		Cost:        0,
	}
}

// Mix beverage with topping
// Return Drink{}
func MixBeverageWithTopping(cup Cup, topping toppings.Topping) Drink {
	cupDescription := cup.GetDescription() + " with topping " + topping.GetDescription()
	cupPrice := cup.GetPrice() + topping.GetPrice()

	return Drink{
		Description: cupDescription,
		Cost:        cupPrice,
	}
}
