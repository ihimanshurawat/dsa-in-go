package tripletinsortedarray

import "slices"

func TripletInSortedArray(nums []int, value int) bool {
	slices.Sort(nums)
	for i := 0; i < len(nums)-2; i++ {
		currentItem := nums[i]
		remainingSum := value - currentItem
		if isPair(nums, i+1, remainingSum) {
			return true
		}
	}
	return false
}


func isPair(nums []int, startingIndex int, value int) bool {
	start := startingIndex
	end := len(nums) - 1

	for start < end {
		if nums[start] + nums[end] == value {
			return true
		} else if nums[start] + nums[end] > value {
			end--;
		} else {
			start++;
		}
	}
	return false;
}
