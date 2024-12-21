package selectionsort

import "math"

//Find Minimum and Place it at its correct location in each pass.
func SelectionSort(nums []int) []int {
	for i:=0;i<len(nums); i++ {
		minValue := math.MaxInt
		minIndex := -1
		for j:=i; j<len(nums); j++ {
			if minValue > nums[j] {
				minValue = nums[j]
				minIndex = j
			} 
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}