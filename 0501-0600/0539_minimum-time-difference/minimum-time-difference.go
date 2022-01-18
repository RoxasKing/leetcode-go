package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Hash
// Sorting

func findMinDifference(timePoints []string) int {
	set := map[int]struct{}{}
	for _, t := range timePoints {
		h, m := int(t[0]-'0')*10+int(t[1]-'0'), int(t[3]-'0')*10+int(t[4]-'0')
		num := h*60 + m
		if _, ok := set[num]; ok {
			return 0
		} else {
			set[num] = struct{}{}
		}
	}
	n := len(set)
	arr := make([]int, 0, n)
	for num := range set {
		arr = append(arr, num)
	}
	sort.Ints(arr)
	max := 24 * 60
	out := max
	for i := 0; i < n; i++ {
		out = Min(out, (arr[(i+1)%n]-arr[i]+max)%max)
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
