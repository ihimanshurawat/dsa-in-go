package maxconsecutiveones

import (
	"testing"
)

func TestMaxOnes(t *testing.T) {
	testInput := []struct{
		inputData []int
		expectedResult int
	} {
		{
			[]int{0,1,1,1,0,1},
			3,
		},
		{
			[]int{0,0,0,0,0},
			0,
		},
		{
			[]int{1,1,1,1,1},
			5,
		},
		{
			[]int{0,0,0,1,0,0,1,1,1,1,0,0,0,1,1,1,1,1,1,1},
			7,
		},
	}

	for index, value := range testInput {
		actualResult := MaxOnes(value.inputData)

		if value.expectedResult != actualResult {
			t.Fatalf("Test #%d Failed. Expected Result %v Actual Result %v", index, value.expectedResult, actualResult)
		}
	}
}