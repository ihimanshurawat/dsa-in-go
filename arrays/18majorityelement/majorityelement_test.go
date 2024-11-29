package majorityelement

import "testing"

func TestMajorityElement(t *testing.T) {

	testInput := []struct {
		inputData      []int
		expectedResult int
	}{
		{
			[]int{1, 2, 3, 4, 4, 4, 4},
			4,
		},
		{
			[]int{1,1},
			1,
		},
	}

	for index, value := range testInput {
		actualResult := MajorityElement(value.inputData)
		if value.expectedResult != actualResult {
			t.Fatalf("Test #%d Failed, Actual %v and Expected %v", index, actualResult, value.expectedResult)
		}
	}
}
