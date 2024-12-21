

package searchinfinitearray

func SearchInfiniteArray(nums []int, value int) int {
	if nums[0] == value {
		return 0
	}
	i := 1

	for nums[i] < value {
		i = i * 2
	}

	return binarySearch(nums, value, i/2+1, i-1)
}

func binarySearch(nums []int, value int, start int, end int) int {
	for start < end {
		mid := (start + end) / 2
		if nums[mid] == value {
			return mid
		} else if nums[mid] > value {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}
