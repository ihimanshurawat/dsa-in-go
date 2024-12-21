package searchinsortedrotatedarray

func SearchInSortedRotatedArray(nums []int, value int) int {
	start := 0
	end := len(nums)

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] == value {
			return mid
		} else if nums[start] < nums[mid] {
			if nums[start] <= value && value < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] < value && value <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
