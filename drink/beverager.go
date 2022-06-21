package drink

// Beverager struct
type Beverager interface {
	SetCost()
	SetDescription()
	GetCost() float32
	GetBeverager() Drink
}

type Drink struct {
	Description string
	Cost        float32
}

func GetDrink(drinkType int) Beverager {
	switch drinkType {
	case 1:
		return &DarkRoast{}
	case 2:
		return &Milk{}
	case 3:
		return &IceCream{}
	default:
		return nil
	}
}
