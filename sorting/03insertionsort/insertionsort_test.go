package insertionsort

import (
	"slices"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	actualResult := InserstionSort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	expectedResult := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if !slices.Equal(actualResult, expectedResult) {
		t.Fatalf("Test #0 Failed. Actual Result %v and Expected Result %v", actualResult, expectedResult)
	}
}
