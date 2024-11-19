package maximumcircularsumsubarray

import "testing"

func TestMaximumCircularSumSubArray(t *testing.T) {

	testInput := []struct {
		inputData      []int
		expectedResult int
	}{
		{
			[]int{8, -4, 3, -5, 4},
			12,
		},
	}

	for index, value := range testInput {
		actualResult := MaximumCircularSumEfficient(value.inputData)

		if actualResult != value.expectedResult {
			t.Fatalf("Test #%d Failed. Actual Result %v and Expected Result %v", index, actualResult, value.expectedResult)
		}
	}
}
