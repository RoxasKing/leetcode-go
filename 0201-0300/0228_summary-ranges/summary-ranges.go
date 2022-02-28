package main

import "strconv"

// Difficulty:
// Easy

func summaryRanges(nums []int) []string {
	n := len(nums)
	if n == 0 {
		return nil
	}
	l, r := 0, 1
	out := []string{}
	add := func() {
		str := strconv.Itoa(nums[l])
		if l != r-1 {
			str += "->" + strconv.Itoa(nums[r-1])
		}
		out = append(out, str)
	}
	for ; r < n; r++ {
		if nums[r-1]+1 != nums[r] {
			add()
			l = r
		}
	}
	add()
	return out
}
