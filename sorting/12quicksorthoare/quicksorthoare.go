package quicksorthoare

func HQuickSort(nums []int, low int, high int) {
	if low < high {
		partition := partition(nums, low, high)
		HQuickSort(nums, low, partition)
		HQuickSort(nums, partition + 1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[low]
	i, j := low - 1, high + 1

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
				break;
			}
		}

		if i >= j {
			return j
		}
		
		nums[i], nums[j] = nums[j], nums[i]
	}
}
