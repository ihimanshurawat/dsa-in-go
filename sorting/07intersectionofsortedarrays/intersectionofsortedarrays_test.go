package intersectionofsortedarrays

import (
	"reflect"
	"testing"
)

func TestInteresectionOfSortedArrays(t *testing.T) {
	actualResult := IntersectionOfSortedArrays([]int{10, 20, 20, 40, 60}, []int{2,20,20,20})
	expectedResult := []int{20}

	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("Test #0 Failed. Expected %v, got %v", expectedResult, actualResult)
	}
}
