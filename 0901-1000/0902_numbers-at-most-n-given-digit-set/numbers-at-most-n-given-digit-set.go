package main

import "sort"

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

func atMostNGivenDigitSet(digits []string, n int) int {
	length := 0
	for x := n; x > 0; x /= 10 {
		length++
	}
	nums := make([]int, length)
	for i := 0; i < length; i++ {
		nums[i] = n % 10
		n /= 10
	}
	k := len(digits)
	out := 0
	cnts := make([]int, length)
	cnts[0] = 1
	for i := 1; i < length; i++ {
		cnts[i] = cnts[i-1] * k
		out += cnts[i]
	}
	arrs := make([]int, k)
	for i := range digits {
		arrs[i] = int(digits[i][0] - '0')
	}
	sort.Ints(arrs)
	for i := length - 1; i >= 0; i-- {
		j := 0
		for ; j < k; j++ {
			if arrs[j] >= nums[i] {
				break
			}
			out += cnts[i]
		}
		if j == k || arrs[j] != nums[i] {
			break
		}
		if i == 0 {
			out++
		}
	}
	return out
}
