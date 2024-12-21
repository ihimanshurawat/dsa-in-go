package kthsmallestelement

func KthSmallestElement(nums []int, low int, high int, k int) int {
	
	for low <= high {
		p := partition(nums, low, high)

		if p == k-1 {
			return nums[p]
		} else if p > k-1 {
			high = p - 1
		} else {
			low = p + 1
		}
	}
	return -1
}

func partition(nums []int, low int, high int) int {
	pivot := nums[high]
	i := low - 1

	for j := 0; j <= high-1; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}
