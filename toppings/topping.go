package toppings

type Topping interface {
	GetPrice() float32
	GetDescription() string
}

type Top struct {
	Description string
	Cost        float32
}

func NewTopping() Topping {
	return &Top{
		Description: "",
		Cost:        0,
	}
}

func (top *Top) GetPrice() float32 {
	return 0
}

func (top *Top) GetDescription() string {
	return ""
}
