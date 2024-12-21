package unionofsortedarrays

func UnionOfSortedArrays(nums1 []int, nums2 []int) []int {
	var result []int

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if i > 0 && nums1[i] == nums1[i-1] {
			i++
			continue
		}

		if j > 0 && nums2[j] == nums2[j-1] {
			j++
			continue
		}

		if nums1[i] < nums2[j] {
			result = append(result, nums1[i])
			i++
		} else if nums1[i] > nums2[j] {
			result = append(result, nums2[j])
			j++
		} else {
			result = append(result, nums1[i])
			i++
			j++
		}
	}

	for i < len(nums1) {
		if i > 0 && nums1[i] == nums1[i-1] {
			i++
		} else {
			result = append(result, nums1[i])
			i++
		}
	}

	for j < len(nums2) {
		if j > 0 && nums2[j] == nums2[j-1] {
			j++
		} else {
			result = append(result, nums2[j])
			j++
		}
		
	}
	return result
}
