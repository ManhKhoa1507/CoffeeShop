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
	cup       Cup
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

func NewDirector(b Beverager) *Director {
	return &Director{
		beverager: b,
	}
}

func (d *Director) SetBuilder(b Beverager) {
	d.beverager = b
}

func (d *Director) BuildBeverager() Drink {
	d.beverager.SetCost()
	d.beverager.SetDescription()
	return d.beverager.GetBeverager()
}
