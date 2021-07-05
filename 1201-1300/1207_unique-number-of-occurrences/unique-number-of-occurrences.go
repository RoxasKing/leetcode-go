package main

import "sort"

func uniqueOccurrences(arr []int) bool {
	count := map[int]int{}
	nums := []int{}
	for _, a := range arr {
		if _, ok := count[a]; !ok {
			nums = append(nums, a)
		}
		count[a]++
	}
	sort.Slice(nums, func(i, j int) bool { return count[nums[i]] < count[nums[j]] })
	for i := 1; i < len(nums); i++ {
		if count[nums[i]] == count[nums[i-1]] {
			return false
		}
	}
	return true
}

func uniqueOccurrences2(arr []int) bool {
	valCount := make([]int, 2001)
	for _, a := range arr {
		valCount[a+1000]++
	}
	countMark := make([]int, 1001)
	for _, count := range valCount {
		if count == 0 {
			continue
		}
		if countMark[count] != 0 {
			return false
		}
		countMark[count] = 1
	}
	return true
}
