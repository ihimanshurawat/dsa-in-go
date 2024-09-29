package reversearray

import (
	"slices"
	"testing"
)

func TestReverseArray(t *testing.T) {
	testInput := [] struct {
		input []int
		expectedResult []int
	} {
		{
			[]int{1,2,3}, 
			[]int{3,2,1},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
	}

	for index, value := range testInput {
		actualResult := ReverseArray(value.input)

		if !slices.Equal(actualResult, value.expectedResult) {
			t.Fatalf("Test failed for Input Index #%d", index)
		}
	}
}