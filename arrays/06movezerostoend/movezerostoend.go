package movezerostoend

func MoveZerosToEnd(arr []int) []int {
	nonZeroIndex := -1;
	for i:=0; i<len(arr); i++ {
		if arr[i] != 0 {
			nonZeroIndex++
			temp := arr[i]
			arr[i] = arr[nonZeroIndex]
			arr[nonZeroIndex] = temp;
		}
 	}
	return arr;
}