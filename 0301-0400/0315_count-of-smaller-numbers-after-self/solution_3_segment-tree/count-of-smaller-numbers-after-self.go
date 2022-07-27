package main

import "sort"

// Difficulty:
// Hard

// Tags:
// Segment Tree

func countSmaller(nums []int) []int {
	set := map[int]struct{}{}
	arr := []int{}
	for _, x := range nums {
		if _, ok := set[x]; !ok {
			set[x] = struct{}{}
			arr = append(arr, x)
		}
	}
	sort.Ints(arr)
	h := map[int]int{}
	for i, v := range arr {
		h[v] = i
	}
	n := len(arr)
	f := make([]int, n*4)
	s, t := 0, n-1
	var add func(i, s, t, x int)
	add = func(i, s, t, x int) {
		if x < s || t < x {
			return
		}
		if s == t {
			f[i]++
			return
		}
		m := (s + t) / 2
		add(i*2, s, m, x)
		add(i*2+1, m+1, t, x)
		f[i] = f[i*2] + f[i*2+1]
	}
	var query func(i, s, t, r int) int
	query = func(i, s, t, r int) int {
		if r < s {
			return 0
		}
		if t <= r {
			return f[i]
		}
		m := (s + t) / 2
		return query(i*2, s, m, r) + query(i*2+1, m+1, t, r)
	}
	o := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		x := h[nums[i]]
		o[i] = query(1, s, t, x-1)
		add(1, s, t, x)
	}
	return o
}
