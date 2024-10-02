package removeduplicates

/*
RemoveDuplicates is a function in Go to remove duplicates from a sorted array. 
*/
func RemoveDuplicates(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	
	var uniqueIndex int = 0;
	for i:=1; i<len(arr); i++ {
		if arr[i-1] != arr[i] {
			uniqueIndex++;
			arr[uniqueIndex] = arr[i]
		}
	}
	return arr[:uniqueIndex + 1]
}