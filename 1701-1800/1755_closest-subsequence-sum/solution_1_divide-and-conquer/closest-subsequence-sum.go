package main

import (
	"sort"
)

// Tags:
// Divide-and-Conquer Algorithm

func minAbsDifference(nums []int, goal int) int {
	out := Min(Abs(goal), Abs(nums[0]-goal))
	n := len(nums)
	if n == 1 {
		return out
	}

	m := n >> 1
	lp := partition(nums[:m])
	rp := partition(nums[m:])
	arr := make([]int, 0, len(lp))
	for num := range lp {
		out = Min(out, Abs(num-goal))
		arr = append(arr, num)
	}
	sort.Ints(arr)

	for num := range rp {
		out = Min(out, Abs(num-goal))
		i := sort.Search(len(arr), func(i int) bool { return num+arr[i] >= goal })
		if i < len(arr) {
			out = Min(out, Abs(arr[i]+num-goal))
		}
		if i > 0 {
			out = Min(out, Abs(arr[i-1]+num-goal))
		}
	}
	return out
}

func partition(nums []int) map[int]struct{} {
	n := len(nums)
	if n == 1 {
		return map[int]struct{}{nums[0]: {}}
	}
	m := n >> 1
	lp := partition(nums[:m])
	rp := partition(nums[m:])
	set := map[int]struct{}{}
	for num1 := range lp {
		set[num1] = struct{}{}
		for num2 := range rp {
			set[num2] = struct{}{}
			set[num1+num2] = struct{}{}
		}
	}
	return set
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
