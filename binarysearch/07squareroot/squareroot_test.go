package squareroot

import "testing"

func TestSquareRoot(t *testing.T) {
	actualResult := SquareRoot(17)
	expectedResult := 4

	if actualResult != expectedResult {
		t.Fatalf("Test #0 Failed. Actual Result %v and Expected Result %v", actualResult, expectedResult)
	}
}
