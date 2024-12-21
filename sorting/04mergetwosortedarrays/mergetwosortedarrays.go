package mergetwosortedarrays

func MergeTwoSortedArrays(first []int, second []int) []int {
	var mergedArray []int 

	i:=0;j:=0;
	for i<len(first) && j<len(second) {
		if first[i] <= second[j] {
			mergedArray = append(mergedArray, first[i])
			i++
		} else {
			mergedArray = append(mergedArray, second[j])
			j++
		}
	}

	for i<len(first) {
		mergedArray = append(mergedArray, first[i])
		i++
	}

	for j<len(second) {
		mergedArray = append(mergedArray, second[j])
		j++
	}

	return mergedArray
}
