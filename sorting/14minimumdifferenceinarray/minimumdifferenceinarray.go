package minimumdifferenceinarray

import (
	"math"
	"slices"
)

func MinimumDifferenceInArray(nums []int) int {
	slices.Sort(nums)
	result := math.MaxFloat64
	for i:=1; i<len(nums); i++ {
		result = math.Min(float64(nums[i] - nums[i-1]), result)
	}
	return int(result)
}