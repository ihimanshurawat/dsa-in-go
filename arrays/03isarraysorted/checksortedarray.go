package isarraysorted


/*
Function to check if the array is sorted in Go. 
*/
func IsArraySorted(arr []int) bool {
	if len(arr) < 2 {
		return true
	}

	for i:=1; i<len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}