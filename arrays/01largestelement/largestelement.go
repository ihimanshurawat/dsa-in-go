package largestelement

/*
LargestElement function will return the largest element from the input array.
 */
func LargestElement(arr []int) int {
	var largestElement int = -1;

	for i:=0; i<len(arr); i++ {
		if largestElement < arr[i] {
			largestElement = arr[i];
		}
	}
	
	return largestElement
}
