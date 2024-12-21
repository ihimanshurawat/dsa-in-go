package mergefunction

import (
	"reflect"
	"testing"
)

func TestMergeFunction(t *testing.T) {
	// Input data and expected result
	input := []int{6,7,8,9,10,1, 2, 3,4,5}
	expected := []int{1,2,3,4,5,6,7,8,9,10} // Expected result after merging indices [2:6] and [6:9]

	// Indices for the merge
	low, mid, high := 0, 4, 9

	// Perform the merge
	MergeFunction(input, low, mid, high)

	// Verify the result
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("MergeFunction() failed. Expected %v, got %v", expected, input)
	}
}