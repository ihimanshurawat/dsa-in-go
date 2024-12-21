package countonesinsortedarray

import "testing"

func TestCountOnesInSortedArray(t *testing.T) {
	actualResult := CountOnesInSortedArray([]int {0,1,1})
	if actualResult != 2 {
		t.Fatalf("Test #0 Failed. Actual Result %v and Expected Result %v", actualResult, 2)
	}
}
