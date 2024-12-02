package indexoffirstoccurrence

import "testing"

func TestIndexOfFirstOccurence(t *testing.T) {
	testInput := []struct {
		inputData      []int
		value          int
		expectedResult int
	}{
		{
			[]int{9, 10, 10, 10, 11, 12, 13, 15, 16},
			10,
			1,
		},
		{
			[]int{9, 10, 10, 10, 11, 12, 13, 15, 16},
			16,
			8,
		},
	}

	for index, value := range testInput {
		actualResult := IndexOfFirstOccurence(value.inputData, value.value)

		if value.expectedResult != actualResult {
			t.Fatalf("Test $%dM Failed. Actual Result %v, Expected Result %v", index, actualResult, value.expectedResult)
		}
	}
}
