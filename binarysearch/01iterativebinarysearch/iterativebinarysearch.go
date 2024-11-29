package iterativebinarysearch

//[]int{1,2,3,4,5,6,7,8,9,10} and 4
func IterativeBinarySearch(nums []int, value int) int {
	start := 0 
	end := len(nums) - 1 //9 

	// 0 < 9 true // 0 < 3 true // 2 < 3
	for start <= end {
		mid := (start + end) / 2 // 9/2 = 4 // 3 / 2 = 1 // 5 / 2 = 2
		if nums[mid] == value { // nums[4] 5 != 4 // nums[1] 2 != 4 // nums[2] 3 != 4
			return mid
		} else if nums[mid] > value {
			end = mid - 1 //end = 4 - 1 = 3 // X // X
		} else {
			start = mid + 1 // X // start =  2 // start = 3
		}
	}
	return -1;
}
