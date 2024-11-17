package longestevenoddsubarray

import "testing"

func TestLongestEvenOddSubArray(t *testing.T) {
	testInput := []struct {
		inputData []int
		expectedResult int
	} {
		{
			[]int{1,2,3,4,5,6},
			6,
		},
		{
			[]int{1,1,1,1,1,1},
			1,
		},
	}

	for index, value := range testInput {
		actualResult := LongestEvenOddSubArray(value.inputData)

		if value.expectedResult != actualResult {
			t.Fatalf("Test #%d Failed. Expected Result %v Actual Result %v", index, value.expectedResult, actualResult)
		}
	}
}