package movezerostoend

import (
	"slices"
	"testing"
)

func TestMoveZerosToEnd(t *testing.T) {
	testInput := []struct {
		input []int
		expectedResult []int
	} {
		{
			[]int{1,2,0,0,2,0,1},
			[]int{1,2,2,1,0,0,0},
		},
		{
			[]int{0},
			[]int{0},
		},
		{
			[]int{},
			[]int{},
		},
	}

	for index, value := range testInput {
		actualResult := MoveZerosToEnd(value.input)
		if !slices.Equal(value.expectedResult, actualResult) {
			t.Fatalf("Test #%d failed. Expected Result %v and Actual Result %v", index, value.expectedResult, actualResult)
		}
	}
}