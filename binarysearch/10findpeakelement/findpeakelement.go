package findpeakelement

func FindPeakElement(nums []int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2

		if mid-1 < 0 || nums[mid-1] < nums[mid] && mid+1 >= len(nums) || nums[mid+1] < nums[mid] {
			return nums[mid]
		} else if mid > 0 && nums[mid-1] > nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
