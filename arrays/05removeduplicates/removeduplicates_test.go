package removeduplicates

import (
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testInput := [] struct {
		input []int
		expectedResult []int	
	} {
		{
			[]int{1,1,2,2,3,4,5,5,5,6,6,6,7},
			[]int{1,2,3,4,5,6,7},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1,2,3,4,5},
			[]int{1,2,3,4,5},
		},
		{
			[]int{1,1,1,1},
			[]int{1},
		},
	}

	for index, value := range testInput {
		actualResult := RemoveDuplicates(value.input)

		if !slices.Equal(actualResult, value.expectedResult) {
			t.Fatalf("Failed test execution for Input #%d with Actual Result %v and Expected Result %v", index, actualResult, value.expectedResult)
		}

	}
}