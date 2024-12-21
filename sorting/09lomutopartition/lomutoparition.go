package lomutopartition

// Assuming Pivot is last Element of the input array nums
func LomutoParition(nums []int, low int, high int) int {
	i := low - 1
	pivot := nums[high]
	for j := low; j < high-low+1; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}
