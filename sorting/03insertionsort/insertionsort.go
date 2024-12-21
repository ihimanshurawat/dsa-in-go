package insertionsort

func InserstionSort(nums []int) []int {
	//Scan all elements
	for i:=1; i<len(nums); i++ {
		temp := nums[i]
		for j:=i-1; j>=0; j-- {
			if nums[j] > temp {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			} else {
				j++
				nums[j] = temp
			}
		}
	}

	return nums
}