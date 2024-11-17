package longestevenoddsubarray

func LongestEvenOddSubArray(nums []int) int {
	count := 1;
	res := 1;
	for i:=1; i<len(nums); i++ {
		if nums[i] % 2 != nums[i-1] % 2 {
			count++;
		} else {
			count = 1;
		}
		if count > res {
			res = count;
		}
	}
	return res;
}