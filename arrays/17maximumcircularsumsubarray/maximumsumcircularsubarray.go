package maximumcircularsumsubarray

func MaximumCircularSum(nums []int) int {
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		curr := nums[i]
		currMax := nums[i]
		for j := 1; j < len(nums); j++ {
			index := (i + j) % len(nums)
			curr += nums[index]
			if curr > currMax {
				currMax = curr
			}
		}
		if currMax > res {
			res = currMax
		}
	}

	return res
}

func MaximumCircularSumEfficient(nums []int) int {
	normalSum := maxSumUsingKadane(nums)

	if normalSum < 0 {
		return normalSum
	}

	arraySum := 0;

	for i:=0; i<len(nums); i++ {
		arraySum += nums[i]
	}

	circularSum := arraySum - minSumUsingKadane(nums);

	if circularSum > normalSum {
		return circularSum
	} else {
		return normalSum
	}
}

func maxSumUsingKadane(nums []int) int {
	res := nums[0]
	maximumSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if maximumSum+nums[i] > nums[i] {
			maximumSum += nums[i]
		} else {
			maximumSum = nums[i]
		}

		if res < maximumSum {
			res = maximumSum
		}
	}
	return res
}

func minSumUsingKadane(nums []int) int {
	res := nums[0]
	minimumSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if minimumSum+nums[i] < nums[i] {
			minimumSum += nums[i]
		} else {
			minimumSum = nums[i]
		}

		if res > minimumSum {
			res = minimumSum
		}
	}

	return res
}
