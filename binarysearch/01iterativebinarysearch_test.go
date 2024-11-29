package iterativebinarysearch

import "testing"

func TestIterativeBinarySearch(t *testing.T) {
	testInput := []struct {
		inputData      []int
		searchValue    int
		expectedResult int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			4,
			3,
		},
	}

	for index, value := range testInput {
		actualResult := IterativeBinarySearch(value.inputData, value.searchValue)

		if value.expectedResult != actualResult {
			t.Fatalf("Test #%d Failed. Actual %v and Expected %v", index, actualResult, value.expectedResult)
		}
	}
}
