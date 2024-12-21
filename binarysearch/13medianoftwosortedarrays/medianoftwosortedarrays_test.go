package medianoftwosortedarrays

import (
	"testing"
)

func TestMedianOfTwoSortedArrays(t *testing.T) {
	//[]int{10,20,30,40,50}, []int{5,15,25,35,45,55,65,75,85}
	actualResult := MedianOfTwoSortedArrays([]int{1, 2, 3, 4, 5, 6}, []int{10, 20, 30, 40, 50})
	expectedResult := 6

	if actualResult != float64(expectedResult) {
		t.Fatalf("Test #0 Failed. Actual Result %v and Expected Result %v", actualResult, expectedResult)
	}
}
