package leftrotatebyd

/*
Function to left rotate array by D times in Go Lang. 
*/
func LeftRotateByD(arr []int, d int) []int {
	reverse(arr, 0, d-1)
	reverse(arr, d, len(arr) - 1)
	reverse(arr, 0, len(arr) - 1)
	return arr;
}

func reverse(arr []int, low int, high int) {
	for low < high {
		arr[low], arr[high] = arr[high], arr[low]
		low++
		high--
	}
}