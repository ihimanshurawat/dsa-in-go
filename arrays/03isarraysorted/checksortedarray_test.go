package isarraysorted

import "testing"

func TestIsArraySorted(t *testing.T) {
	testInput := []struct{
		input []int
		expectedResult bool
	} {
		{[]int{1,3,4,5,3}, false},
		{[]int{1,2,3,4,5}, true},
		{[]int{}, true},
		{[]int{1}, true},
	}

	for index, value := range testInput {
		actualResult := IsArraySorted(value.input)

		if actualResult != value.expectedResult {
			t.Fatalf("Test failed for #%d", index)
		}
	}
}

