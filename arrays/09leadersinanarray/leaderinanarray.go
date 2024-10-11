package leadersinanarray

func LeadersInArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	
	result := []int{}

	//Last element of the array is always the Leader
	result = append(result, nums[len(nums) - 1])
	max := nums[len(nums) - 1]

	for i:=len(nums)-2; i>=0;i-- {
		if nums[i] > max {
			max = nums[i];
			result = append(result, nums[i])
		}
	}
	return result;
}