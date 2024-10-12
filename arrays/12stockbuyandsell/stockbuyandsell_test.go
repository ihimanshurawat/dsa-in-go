package stockbuyandsell

import "testing"

func TestStockBuyAndSell(t *testing.T) {
	testInput := []struct {
		input []int
		expectedResult int
	} {
		{
			[]int{1,5,3,8,12},
			13,
		},
		{
			[]int{5,5,5},
			0,
		},
		{
			[]int{5,10,15},
			10,
		},
	}

	for index, value := range testInput {
		actualResult := StockBuyAndSell(value.input)

		if actualResult != value.expectedResult {
			t.Fatalf("Test #%d Failed. Expected Result %v Actual Result %v", index, value.expectedResult, actualResult)
		}
	}
}