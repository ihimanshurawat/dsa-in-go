package leftrotatebyone

func LeftRotateByOne(arr []int) []int {
	if len(arr) < 2 {
		return arr;
	}
	first := arr[0]

	for i:=1; i<len(arr); i++ {
		arr[i-1] = arr[i]
	}
	arr[len(arr) - 1] = first;
	return arr;
}