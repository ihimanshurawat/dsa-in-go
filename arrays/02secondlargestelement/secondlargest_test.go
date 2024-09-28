package secondlargest

import (
	"testing"
)

func TestSecondLargest(t *testing.T) {
	testInput := []struct {
		input          []int
		expectedResult int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, -1},
	}

	for index, value := range testInput {
		actualResult := SecondLargest(value.input)

		if actualResult != value.expectedResult {
			t.Fatalf("Test Run #%d Failed. Expected Value %v != Actual Value %v", index, value.expectedResult, actualResult)
		}
	}
}
