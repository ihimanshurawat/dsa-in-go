package majorityelement

import "slices"

//If an element occurres > size/2 times then it is a majority element
func MajorityElement(nums []int) int {
	//Sorting input array using variation of quicksort so O(nlogn)
	slices.Sort(nums)
	count := 1
	for i:=1; i<len(nums); i++ {
		if nums[i-1] == nums[i] {
			count++
		} else {
			count = 1
		}
		if count > len(nums) / 2 {
			return nums[i]
		}
	}
	return -1;
}