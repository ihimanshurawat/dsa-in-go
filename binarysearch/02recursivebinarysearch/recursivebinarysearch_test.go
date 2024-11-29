package recursivebinarysearch

import (
	"testing"
)

func TestRecursiveBinarySearch(t *testing.T) {
	testInput := []struct {
		inputData      []int
		value          int
		expectedResult int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			4,
			3,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11},
			10,
			-1,
		},
	}

	for index, value := range testInput {
		actualResult := RecursiveBinarySearch(value.inputData, 0, len(value.inputData)-1, value.value)

		if value.expectedResult != actualResult {
			t.Fatalf("Test #%d Failed. Actual %v and Expected %v", index, actualResult, value.expectedResult)
		}
	}
}
