package quicksortlomuto

func LQuickSort(nums []int, low int, high int) {
	if low < high {
		partition := partition(nums, low, high)
		LQuickSort(nums, low, partition-1)
		LQuickSort(nums, partition+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}
