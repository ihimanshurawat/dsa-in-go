package searchinsortedrotatedarray

import "testing"

func TestSearchInSortedRotatedArray(t *testing.T) {
	actualResult := SearchInSortedRotatedArray([]int{10, 20, 40, 60, 5, 8}, 5)
	expectedResult := 4

	if actualResult != expectedResult {
		t.Fatalf("Test #0 Failed. Actual Result %v and Expected Result %v", actualResult, expectedResult)
	}
}
