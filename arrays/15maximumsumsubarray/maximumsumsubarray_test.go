package maximumsumsubarray

import "testing"

func TestMaximumSumSubArray(t *testing.T) {
	testInput := []struct{
		inputData []int
		expectedResult int
	} {
		{
			[]int{1,2,3},
			6,
		},
		{
			[]int{-1,-2,3},
			3,
		},
	}

	for index, value := range testInput {
		actualResult := MaximumSumSubArray(value.inputData)

		if value.expectedResult != actualResult {
			t.Fatalf("Test #%d Failed. Expected Value %v and Actual Value %v", index, value.expectedResult, actualResult)
		}
	}
}