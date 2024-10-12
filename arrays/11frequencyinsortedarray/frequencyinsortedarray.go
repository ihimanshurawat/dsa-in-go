package main

import (
	"fmt"
)

func FrequencyInSortedArray(nums []int) {
	freq := 1
	i := 1 
	for i < len(nums) {
		for i < len(nums) && nums[i] == nums[i-1] {
			freq++;
			i++;
		}
		fmt.Printf("Element %v has Frequency %v \n", nums[i-1], freq)
		i++;
		freq = 1;
	}

	if len(nums) == 1 || nums[len(nums) - 1] != nums[len(nums) - 2] {
		fmt.Printf("Element %v has Frequency 1 \n", nums[len(nums) - 1])
	}
}
