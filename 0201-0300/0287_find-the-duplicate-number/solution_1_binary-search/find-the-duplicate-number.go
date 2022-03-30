package main

// Difficulty:
// Medium

// Tags:
// Binary Search

func findDuplicate(nums []int) int {
	l, r := 1, len(nums)-1
	for l < r {
		m := (l + r) >> 1
		cnt := 0
		for _, num := range nums {
			if num <= m {
				cnt++
			}
		}
		if cnt <= m {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
