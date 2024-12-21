package kthsmallestelement

import (
	"reflect"
	"testing"
)

func TestKthSmallestElement(t *testing.T) {
	input := []int{6, 5, 8, 9, 10, 11, 12, 13, 4, 1, 3, 1, 1, 2, 3, 4, 5}
	actualResult := KthSmallestElement(input, 0, len(input)-1, 2)
	expectedResult := 1

	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("Test #0 Failed. Expected %v, got %v", expectedResult, actualResult)
	}
}
