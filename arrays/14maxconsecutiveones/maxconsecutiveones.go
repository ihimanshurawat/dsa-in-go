package maxconsecutiveones

func MaxOnes(nums []int) int {
	curr := 0
	res := 0

	for i:=0;i<len(nums); i++ {
		if nums[i] == 0 {
			curr = 0;
		} else {
			curr++;
			if curr > res {
				res = curr;
			}
		}
	}

	return res
}