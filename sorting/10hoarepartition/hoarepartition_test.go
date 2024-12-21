package hoarepartition

import (
	"reflect"
	"testing"
)

func TestHoareParition(t *testing.T) {
	input := []int{5, 3, 8, 4, 2, 7, 1,10}
	HoarePartition(input, 0, len(input) - 1)
	expectedResult := []int{1,3,4,2,5,8,7,10}

	if !reflect.DeepEqual(input, expectedResult) {
		t.Errorf("Test #0 Failed. Expected %v, got %v", expectedResult, input)
	}

}
