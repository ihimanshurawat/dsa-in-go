package countonesinsortedarray


func CountOnesInSortedArray(nums []int) int {
	lastElementIndex := len(nums) - 1
	if nums[lastElementIndex] != 1 {
		return 0
	}

	start := 0
	end := lastElementIndex

	firstIndex := -1

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] == 0 {
			start = mid + 1
		} else {
			if mid - 1 <= 0 || nums[mid - 1] == 0 {
				firstIndex = mid
				break
			} else {
				end = mid - 1
			}
		}
	}
	return lastElementIndex - firstIndex + 1
}