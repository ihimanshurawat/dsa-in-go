package unionofsortedarrays

import (
	"reflect"
	"testing"
)

func TestUnionOfSortedArrays(t *testing.T) {
	actualResult := UnionOfSortedArrays([]int{1,2,3,4,5}, []int{4,4,4,5,5,5,6,6,6,6,6})
	expectedResult := []int{1,2,3,4,5,6}

	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("Test #0 Failed. Expected %v, got %v", expectedResult, actualResult)
	}
}