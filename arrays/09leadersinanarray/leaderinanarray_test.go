package leadersinanarray

import (
	"slices"
	"testing"
)

func TestLeadersInArray(t *testing.T) {
	testInput := []struct{
		input []int
		expectedResult []int
	} {
		{
			[]int{10,20,30},
			[]int{30},
		},
		{
			[]int{30,20,10},
			[]int{10,20,30},
		},
		{
			[]int{1},
			[]int{1},
		},
	}


	for index, value := range testInput {
		actualResult := LeadersInArray(value.input)

		if !slices.Equal(value.expectedResult, actualResult) {
			t.Fatalf("Test #%d Failed. Expected Result %v and Actual Result %v", index, value.expectedResult, actualResult)
		}
	}
}