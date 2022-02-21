package main

// Difficulty:
// Easy

// Tags:
// Boyer-Moore

func majorityElement(nums []int) int {
	out, cnt := -1, 0
	for _, num := range nums {
		if cnt == 0 {
			out = num
		}
		if out == num {
			cnt++
		} else {
			cnt--
		}
	}
	return out
}
