package main

import "sort"

// Binary Search
func lengthOfLIS(nums []int) int {
	mins := []int{}
	for _, num := range nums {
		i := sort.Search(len(mins), func(i int) bool { return mins[i] >= num })
		if i == len(mins) {
			mins = append(mins, num)
		} else {
			mins[i] = num
		}
	}
	return len(mins)
}
