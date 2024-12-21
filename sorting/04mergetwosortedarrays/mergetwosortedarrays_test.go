package mergetwosortedarrays

import (
	"slices"
	"testing"
)

func TestMergeTwoSortedArrays(t *testing.T) {
	actualResult := MergeTwoSortedArrays([]int{1, 2, 3, 7, 8, 9}, []int{4, 5, 6})
	expectedResult := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !slices.Equal(actualResult, expectedResult) {
		t.Fatalf("Test #0 Failed. Actual Result %v and Expected Result %v", actualResult, expectedResult)
	}
}
