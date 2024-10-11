package leftrotatebyd

import (
	"slices"
	"testing"
)

func TestLeftRotateByD(t *testing.T) {

	testInput := []struct {
		input []int
		d int
		expectedResult []int
	} {
		{
			[]int{1,2,3,4,5},
			2,
			[]int{3,4,5,1,2},
		},
		{
			[]int{1,2,3,4,5},
			0,
			[]int{1,2,3,4,5},
		},
	}

	for index, value := range testInput {

		actualResult := LeftRotateByD(value.input, value.d)

		if !slices.Equal(value.expectedResult, actualResult) {
			t.Fatalf("Test #%d Failed. Expected Result %v and Actual Result %v", index, value.expectedResult, actualResult)
		}
	}

}