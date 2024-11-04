package maximumsumsubarray

func MaximumSumSubArray(nums []int) int {
	res := nums[0]
	maxEnding := nums[0]

	for i:=1; i<len(nums); i++ {
		if maxEnding + nums[i] > nums[i] {
			maxEnding += nums[i]
		} else {
			maxEnding = nums[i]
		}
		if maxEnding > res {
			res = maxEnding
		}
	}
	return res;
}
