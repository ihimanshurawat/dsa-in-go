package maximumdifference

import "testing"

func TestMaximumDifference(t *testing.T) {
	testInput := []struct {
		input          []int
		expectedResult int
	} {
		{
			[]int{1, 2, 3, 4, 5, 6},
			5,
		},
		{
			[]int{1},
			0,
		},
	}

	for index, value := range testInput {
		actualResult := MaximumDifference(value.input)
		if actualResult != value.expectedResult {
			t.Fatalf("Test #%d Failed. Expected Result %v and Actual Result %v", index, value.expectedResult, actualResult)
		}
	}
}
