package drink

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
