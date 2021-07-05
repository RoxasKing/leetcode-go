package main

import "strconv"

func summaryRanges(nums []int) []string {
	out := []string{}
	n := len(nums)
	if n == 0 {
		return out
	}
	l, r := nums[0], nums[0]
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1]+1 {
			if l != r {
				out = append(out, strconv.Itoa(l)+"->"+strconv.Itoa(r))
			} else {
				out = append(out, strconv.Itoa(l))
			}
			l, r = nums[i], nums[i]
		} else {
			r = nums[i]
		}
	}
	if l != r {
		out = append(out, strconv.Itoa(l)+"->"+strconv.Itoa(r))
	} else {
		out = append(out, strconv.Itoa(l))
	}
	return out
}
