package findpeakelement

import "testing"

func TestFindPeakElement(t *testing.T) {
	actualResult := FindPeakElement([]int{6, 7, 8, 20, 12})
	expectedResult := 20

	if actualResult != expectedResult {
		t.Fatalf("Test #0 Failed. Actual Result %v and Expected Result %v", actualResult, expectedResult)
	}
}
