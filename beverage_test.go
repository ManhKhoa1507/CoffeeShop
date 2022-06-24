package main

import (
	"CoffeeBill/drink"
	"testing"
)

func TestBuildDrink(t *testing.T) {
	drinkTests := []struct {
		order                    int
		drinkDescriptionExpected string
		drinkCostExpected        float32
	}{
		{order: 1, drinkDescriptionExpected: "Dark roast", drinkCostExpected: 20},
		{order: 2, drinkDescriptionExpected: "Milk", drinkCostExpected: 10},
		{order: 3, drinkDescriptionExpected: "Ice cream", drinkCostExpected: 15},
	}
	for _, test := range drinkTests {

		drink := drink.BuildDrink(test.order)
		drinkDesciption := drink.Description
		drinkCost := drink.Cost

		if drinkDesciption != test.drinkDescriptionExpected || drinkCost != test.drinkCostExpected {
			t.Errorf("Not equal")
		}
	}
}
