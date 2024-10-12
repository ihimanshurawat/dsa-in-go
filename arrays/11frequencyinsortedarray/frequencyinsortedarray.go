package frequencyinsortedarray

import (
	"fmt"
)

func FrequencyInSortedArray(nums []int) {
	elementIndex := 0
	var newElementIndex int
	for i:=1; i<len(nums); i++ {
		if nums[i] != nums[elementIndex] {
			newElementIndex = i
			fmt.Printf("Element %v has Frequency %d", nums[elementIndex], newElementIndex - elementIndex)
			elementIndex = newElementIndex
		}

		if i == len(nums) - 1 {
			newElementIndex = len(nums)
			fmt.Printf("Element %v has Frequency %d", nums[elementIndex], newElementIndex - elementIndex)
		}
	}
}
