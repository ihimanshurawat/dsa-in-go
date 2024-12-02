package indexoffirstoccurrence

func IndexOfFirstOccurence(nums []int, value int) int {
	start := 0 
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] > value {
			end = mid - 1
		} else if nums[mid] < value {
			start = mid + 1
		} else {
			if mid < 1 || nums[mid - 1] != nums[mid] {
				return mid;
			} else {
				end = mid - 1
			}
		}
	}

	return -1;
}