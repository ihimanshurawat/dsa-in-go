package countoccurrences

import "testing"

func TestCountOccurrences(t *testing.T) {
	testInput := []int{1, 2, 3, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5}
	value := 5
	actualResult := CountOccurrences(testInput, value)
	expectedResult := 4

	if actualResult != expectedResult {
		t.Fatalf("Test #0 Failed. Actual Result %v and Expected Result %v", actualResult, expectedResult)
	}
}
