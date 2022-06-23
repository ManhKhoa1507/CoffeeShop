package drink

// Beverager struct
type Beverager interface {
	SetCost()
	SetDescription()
	GetBeverager() Drink
}

type Drink struct {
	Description string
	Cost        float32
}

type Director struct {
	beverager Beverager
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
		// Not such a beverage
		return nil
	}
}

// Get drink order and build drink by calling director
// Choose drink to build
// After build drink return object drink{description, costs}
func BuildDrink(order int) Drink {
	drinkType := GetDrink(order)
	director := NewDirector(drinkType)
	dR := director.BuildBeverager()
	return dR
}

func NewDirector(b Beverager) *Director {
	return &Director{
		beverager: b,
	}
}

func (d *Director) BuildBeverager() Drink {
	d.beverager.SetCost()
	d.beverager.SetDescription()
	return d.beverager.GetBeverager()
}
