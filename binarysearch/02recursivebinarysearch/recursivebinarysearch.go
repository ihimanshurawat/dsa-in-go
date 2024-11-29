package recursivebinarysearch

func RecursiveBinarySearch(nums []int, start int, end int, value int) int {
	if start > end {
		return -1;
	} 
	mid := (start + end) / 2
	if nums[mid] == value {
		return mid;
	} else if nums[mid] > value {
		return RecursiveBinarySearch(nums, start, mid - 1, value)
	} else {
		return RecursiveBinarySearch(nums, mid + 1, end, value)
	}
}