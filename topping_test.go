package main

import (
	"CoffeeBill/toppings"
	"testing"
)

func TestNewCream(t *testing.T) {
	toppingTests := []struct {
		Description     string
		Cost            float32
		toppingExpected string
		CostExpected    float32
	}{
		{Description: "", Cost: 0, toppingExpected: "cream", CostExpected: 1},
		// {Description: "milk foam", Cost: 2, toppingExpected: "milk foam cream", CostExpected: 3},
		// {Description: "cream", Cost: 1, toppingExpected: "cream cream", CostExpected: 2},
	}
	for _, test := range toppingTests {

		// Create new topping order
		orderTopping := &toppings.Top{
			Description: test.Description,
			Cost:        test.Cost,
		}

		top := toppings.NewCream(orderTopping)
		toppingDescription := top.GetDescription()
		toppingCost := top.GetPrice()

		if toppingDescription != test.toppingExpected || toppingCost != test.CostExpected {
			t.Errorf("Not equal")
		}
	}
}
