package main

import "sort"

// Binary Search

func canBeIncreasing(nums []int) bool {
	n := len(nums)
	arr := []int{}
	for _, num := range nums {
		i := sort.Search(len(arr), func(i int) bool { return arr[i] >= num })
		if i == len(arr) {
			arr = append(arr, num)
		} else {
			arr[i] = num
		}
	}
	return len(arr) >= n-1
}
