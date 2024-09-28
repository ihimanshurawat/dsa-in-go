package secondlargest

/*
SecondLargest function return the second largest element from the input array.
If all the elements are same it return -1
*/
func SecondLargest(arr []int) int {
	var largest int = -1
	var secondLargest int = -1;

	for i:=0; i < len(arr); i++ {

		if largest < arr[i] {
			secondLargest = largest
			largest = arr[i]
		} else if largest != arr[i] {
			if secondLargest < arr[i] {
				secondLargest = arr[i]
			}
		}
	}

	return secondLargest
}