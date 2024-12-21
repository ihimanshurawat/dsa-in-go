package twosum

func TwoSum(nums []int, value int) bool {
	start := 0
	end := len(nums) - 1

	for start < end {
		if nums[start]+nums[end] == value {
			return true
		} else if nums[start]+nums[end] > value {
			end--
		} else {
			start++
		}
	}
	return false
}
