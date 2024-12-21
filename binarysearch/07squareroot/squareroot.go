package squareroot

func SquareRoot(input int) int {
	start := 1
	end := input
	result := -1

	for start <= end {
		mid := (start + end) / 2

		square := mid * mid

		if square == input {
			return mid
		} else if square > input {
			end = mid - 1
		} else {
			start = mid + 1
			result = mid
		}
	}
	return result
}
