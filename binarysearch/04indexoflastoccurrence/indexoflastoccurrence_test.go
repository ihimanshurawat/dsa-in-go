package indexoflastoccurrence

import "testing"

func TestIndexOfLastOccurrence(t *testing.T) {
	testInput := []struct{
		inputData []int
		value int
		expectedResult int
	} {
		{
			[]int{1,2,3,4,5,6,9,9,9,9,9},
			9,
			10,
		},
	}

	for index, value := range testInput {
		actualResult := IndexOfLastOccurence(value.inputData, value.value)
		
		if actualResult != value.expectedResult {
			t.Fatalf("Test #%d Faile. Actual Result %v and Expected Result %v", index, actualResult, value.expectedResult)
		}
	}
}