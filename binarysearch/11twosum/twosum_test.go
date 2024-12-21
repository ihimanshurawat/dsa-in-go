package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	actualResult := TwoSum([]int{2, 5, 8, 12, 30}, 17)
	expectedResult := true
	if actualResult != expectedResult {
		t.Fatalf("Test #0 Failed. Actual Result %v and Expected Result %v", actualResult, expectedResult)
	}
}
