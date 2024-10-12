package trapingrainwater

import "math"

/*
Traping Rainwater Problem in Go Lang
*/
func TrapingRainWater(nums []int) int {

	size := len(nums)
	 //Pre-processing of the Array
	 leftMax := make([]int, size)
	 rightMax := make([]int, size)

	 //Left Max is First Element
	 leftMax[0] = nums[0]

	 for i:=1; i<size; i++ {
		if nums[i] > leftMax[i-1] {
			leftMax[i] = nums[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	 }

	 rightMax[size - 1] = nums[size - 1]

	 for i:=size - 2; i>=0; i-- {
		if nums[i] > rightMax[i+1] {
			rightMax[i] = nums[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	 }

	 water := 0

	 for i:=1; i<len(nums) - 1; i++ {
		maximumSupportedHeight := int(math.Min(float64(leftMax[i]), float64(rightMax[i])))
		water += maximumSupportedHeight - nums[i];	
	 }

	 return water;
}