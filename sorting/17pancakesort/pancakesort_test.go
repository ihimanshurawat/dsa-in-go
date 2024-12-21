package pancakesort

import (
	"reflect"
	"testing"
)

func TestPancakeSort(t *testing.T) {
	actualResult := pancakeSort([]int{3,2,4,1})
	expectedResult := []int{1,2,3,4}
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("Test #0 Failed. Expected %v, got %v", expectedResult, actualResult)
	}
}