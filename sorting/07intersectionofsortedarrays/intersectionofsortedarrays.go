package intersectionofsortedarrays

func IntersectionOfSortedArrays(nums1 []int, nums2 []int) []int {
	i, j := 0, 0

	var result []int

	for i < len(nums1) && j < len(nums2) {
		if i > 0 && nums1[i] == nums1[i-1] {
			i++
			continue
		}
		if nums1[i] > nums2[j] {
			j++;
		} else if nums1[i] < nums2[j] {
			i++;
		} else {
			result = append(result, nums1[i])
			i++
			j++
		}
	}
	return result;
}
