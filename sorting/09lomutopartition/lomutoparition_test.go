package lomutopartition

import (
	"reflect"
	"testing"
)

func TestLomutoPartion(t *testing.T) {
	actualResult := LomutoParition([]int{40, 50, 60, 70, 30}, 0, 4)
	expectedResult := 0

	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("Test #0 Failed. Expected %v, got %v", expectedResult, actualResult)
	}
}
