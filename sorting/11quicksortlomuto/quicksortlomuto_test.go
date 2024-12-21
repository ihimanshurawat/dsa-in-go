package quicksortlomuto

import (
	"reflect"
	"testing"
)

func TestLQuickSort(t *testing.T) {
	input := []int{5, 4, 3, 1, 3, 4, 3, 1, 6, 8, 4, 7, 33, 4, 5, 1, 3, 43, 6, 34, 3, 4, 343, 1, 234, 34, 5, 6, 23, 3}
	LQuickSort(input, 0, len(input)-1)
	expectedResult := []int{1, 1, 1, 1, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 8, 23, 33, 34, 34, 43, 234, 343}

	if !reflect.DeepEqual(input, expectedResult) {
		t.Errorf("Test #0 Failed. Expected %v, got %v", expectedResult, input)
	}
}
