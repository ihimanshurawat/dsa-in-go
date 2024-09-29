package reversearray

func ReverseArray(arr []int) []int {
	var low int = 0
	var high int = len(arr) - 1
	for low < high {
		var temp int = arr[low];
		arr[low] = arr[high];
		arr[high] = temp;
		low++;
		high--;
	}
	return arr;
}