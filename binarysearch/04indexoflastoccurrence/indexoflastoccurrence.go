package indexoflastoccurrence


/* Find the index of the last occurence of a value in sorted input array nums.
*/
func IndexOfLastOccurence(nums []int, value int) int {
	start := 0
	end := len(nums)

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] > value {
			end = mid - 1
		} else if nums[mid] < value {
			start = mid + 1
		} else {
			if mid+1 >= len(nums) || nums[mid+1] != nums[mid] {
				return mid
			} else {
				start = mid + 1
			}
		}
	}

	return -1
}