package maximumdifference

import "math"

/*
Finding Maximum of nums[j] - nums[i] where j > i
*/
func MaximumDifference(nums []int) int {
	var result int
	var minimum int = nums[0]

	for i:=0;i<len(nums); i++ {
		result = int(math.Max(float64(result), float64(nums[i] - minimum))) 
		minimum = int(math.Min(float64(minimum), float64(nums[i])))
	}
	return result;
}