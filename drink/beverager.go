package drink

// Beverager struct
type Beverager interface {
	SetCost()
	SetDescription()
}

type Drink struct {
	Description string
	Cost        float32
}
