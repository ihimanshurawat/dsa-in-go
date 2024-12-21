package hoarepartition

func HoarePartition(nums []int, low int, high int) int {
	pivot := nums[low]
	i, j := low - 1, high + 1

	//Infinite Loop in Go Lang
	for {
		for {
			i++
			if nums[i] >= pivot {
				break
			}
		}

		for {
			j--
			if nums[j] <= pivot {
				break
			}
		}

		if i >= j {
			return j
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
}
