package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	input := []int{10,9,8,7,6,5,4,3,2,1}
	expectedInput := []int{1,2,3,4,5,6,7,8,9,10}
	low, high := 0, 9
	MergeSort(input, low, high)

	if !reflect.DeepEqual(input, expectedInput) {
		t.Errorf("Test #0 Failed. Expected %v, got %v", expectedInput, input)
	}
}