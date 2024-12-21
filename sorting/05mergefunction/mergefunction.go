package mergefunction

func MergeFunction(nums []int, low int, mid int, high int) {
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
