package countoccurrences

/* Ologn function to count occurrence of a number in a sorted array.
*/
func CountOccurrences(nums []int, value int) int {
	firstIndex := firstOccurrenceIndex(nums, value)
	//No Occurrences to Count
	if firstIndex == -1 {
		return 0
	}

	lastIndex := lastOccurrenceIndex(nums, value)

	return (lastIndex - firstIndex) + 1
}

func firstOccurrenceIndex(nums []int, value int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] > value {
			end = mid - 1
		} else if nums[mid] < value {
			start = mid + 1
		} else {
			if mid-1 <= 0 || nums[mid-1] != nums[mid] {
				return mid
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}

func lastOccurrenceIndex(nums []int, value int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] > value {
			end = mid - 1
		} else if nums[mid] < value {
			start = mid + 1
		} else {
			if mid+1 >= len(nums) || nums[mid+1] != nums[mid] {
				return mid
			} else {
				start = mid + 1
			}
		}
	}
	return -1
}
