	package largestelement

	import "testing"

	func TestLargestElement(t *testing.T) {
		tests := []struct {
			input    []int
			expected int
		} {
			{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
			{[]int{0}, 0},
			{[]int{}, -1},
			{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 10},
		}

		for index, value := range tests {
			testResult := LargestElement(value.input) 

			if testResult != value.expected {
				t.Errorf("Test case no.%d failed. Expected %v Actual %v", index, value.expected, testResult)
			}
		}
	}
