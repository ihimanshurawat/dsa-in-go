package sortthreetypeofelements

import (
	"reflect"
	"testing"
)

func TestSortThreeTypeOfElements(t *testing.T) {
	input := []int{2,0,2,1,1,0}
	SortThreeTypeOfElements(input)
	expectedResult := []int{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2}

	if !reflect.DeepEqual(input, expectedResult) {
		t.Errorf("Test #0 Failed. Expected %v, got %v", expectedResult, input)
	}
}
