package sorttwotypeofelements

import (
	"reflect"
	"testing"
)

func TestSortEvenOddElements(t *testing.T) {
	input := []int{9, 9, 9, 9, 8, 8, 8, 8}
	SortEvenOddElement(input)
	expectedResult := []int{8, 8, 8, 8, 9, 9, 9, 9}

	if !reflect.DeepEqual(input, expectedResult) {
		t.Errorf("Test #0 Failed. Expected %v, got %v", expectedResult, input)
	}
}
