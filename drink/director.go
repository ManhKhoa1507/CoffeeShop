package drink

type Director struct {
	beverager Beverager
	cup       Cup
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
