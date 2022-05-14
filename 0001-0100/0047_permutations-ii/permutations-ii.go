package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Backtracking

func permuteUnique(nums []int) [][]int {
	a := []int{}
	freq := [21]int{}
	for _, x := range nums {
		if freq[x+10] == 0 {
			a = append(a, x)
		}
		freq[x+10]++
	}
	n := len(nums)
	c := make([]int, 0, n)
	var o [][]int
	var backtrack func()
	backtrack = func() {
		if len(c) == n {
			t := make([]int, n)
			copy(t, c)
			o = append(o, t)
			return
		}
		for _, x := range a {
			if freq[x+10] == 0 {
				continue
			}
			c = append(c, x)
			freq[x+10]--
			backtrack()
			freq[x+10]++
			c = c[:len(c)-1]
		}
	}
	sort.Ints(a)
	backtrack()
	return o
}
