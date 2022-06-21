package drink

type Cup interface {
	GetPrice() float32
}

type DrinkCup struct {
}

func (drinkCup *DrinkCup) GetPrice() float32 {
	return 0
}
