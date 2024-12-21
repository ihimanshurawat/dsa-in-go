package pancakesort

func pancakeSort(arr []int) []int {
	result := []int{}
	for i := 0; i < len(arr); i++ {
		//Return Result for Sorted Array
		if isSorted(arr) {
			return result
		}

		var max int = -1
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > max {
				max = j
			}
		}
		result = append(result, max)
		reverse(arr, 0, max)
		result = append(result, len(arr) - i)
		reverse(arr, 0, len(arr) - i - 1)
	}
	return arr
}

func reverse(arr []int, start int, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}
