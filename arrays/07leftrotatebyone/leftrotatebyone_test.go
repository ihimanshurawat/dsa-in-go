package leftrotatebyone

import (
	"testing"
	"slices"
)

func TestLeftRotateByOne(t *testing.T) {
	testInput := []struct {
		input []int
		expectedResult []int
	} {
		{
			[]int{1,2,3,4,5},
			[]int{2,3,4,5,1},
		},
		{
			[]int{-1},
			[]int{-1},
		}, 
		{
			[]int{},
			[]int{},
		},
	}

	for index, value := range testInput {
		t.Logf("Starting Test for #%d for input %v \n", index, value.input)
		actualResult := LeftRotateByOne(value.input)

		if !slices.Equal(value.expectedResult, actualResult) {
			t.Fatalf("Failed Test #%d Expected Result %v and Actual Result %v", index, value.expectedResult, actualResult)
		}
		
		t.Logf("Test finished successfully for #%d", index)
	}
}