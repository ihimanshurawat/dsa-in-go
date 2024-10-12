package trapingrainwater

import "testing"

func TestTrapingRainWater(t *testing.T) {

	testInput := []struct {
		input []int
		expectedResult int
	} {
		{
			[]int{3,0,1,2,5},
			6,
		},
		{
			[]int{5,0,6,2,3},
			6,
		},
	}

	for index, value := range testInput {
		
		actualResult := TrapingRainWater(value.input)

		if value.expectedResult != actualResult {
			t.Fatalf("Test #%d Failed. Expected Result %v Actual Result %v", index, value.expectedResult, actualResult)
		}
	}
}
