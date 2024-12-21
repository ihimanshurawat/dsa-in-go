package mergesort

func MergeSort(nums []int, low int, high int) {
	if high > low {
		mid := low + (high-low)/2
		MergeSort(nums, low, mid)
		MergeSort(nums, mid+1, high)
		merge(nums, low, mid, high)
	}
}

func merge(nums []int, low, mid, high int) {
	left := make([]int, mid-low+1)
	right := make([]int, high-mid)

	copy(left, nums[low:mid+1])
	copy(right, nums[mid+1:high+1])

	i, j, k := 0, 0, low

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			nums[k] = left[i]
			i++
			k++
		} else {
			nums[k] = right[j]
			j++
			k++
		}
	}

	for i < len(left) {
		nums[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		nums[k] = right[j]
		j++
		k++
	}
}
