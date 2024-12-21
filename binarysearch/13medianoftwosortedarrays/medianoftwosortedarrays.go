package medianoftwosortedarrays

import "math"

func MedianOfTwoSortedArrays(a1 []int, a2 []int) float64 {
	//Swap data if a1 is greater than a2
	if len(a1) > len(a2) {
		a1, a2 = a2, a1
	}

	start := 0
	end := len(a1)
	for start <= end {
		i1 := (start + end) / 2
		i2 := (len(a1)+len(a2)+1)/2 - i1

		var min1 int
		if i1 == len(a1) {
			min1 = math.MaxInt
		} else {
			min1 = a1[i1]
		}

		var max1 int
		if i1 == 0 {
			max1 = math.MinInt
		} else {
			max1 = a1[i1-1]
		}

		var min2 int
		if i2 == len(a2) {
			min2 = math.MaxInt
		} else {
			min2 = a2[i2]
		}

		var max2 int
		if i2 == 0 {
			max2 = math.MinInt
		} else {
			max2 = a2[i2-1]
		}

		if min1 >= max2 && max1 <= min2 {
			if (len(a1)+len(a2))%2 == 0 {
				return (math.Max(float64(max1), float64(max2)) + math.Min(float64(min1), float64(min2))) / 2
			} else {
				return math.Max(float64(max1), float64(max2))
			}
		} else if max2 > min1 {
			start = i1 + 1
		} else {
			end = i1 - 1
		}

	}
	return -1
}
