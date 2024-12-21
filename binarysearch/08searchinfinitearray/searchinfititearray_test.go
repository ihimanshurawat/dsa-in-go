package searchinfinitearray

import "testing"

func TestSearchInfiniteArray(t *testing.T) {
	actualResult := SearchInfiniteArray([]int{10, 20, 30, 40, 50, 60, 80, 90}, 30)

	expectedResult := 2

	if actualResult != expectedResult {
		t.Fatalf("Test #0 Failed. Actual Result %v and Expected Result %v", actualResult, expectedResult)
	}

}
