package main

import (
	"fmt"
	"sort"
)

// Tags:
// Dynamic Programming
// Two Pointers

func minAbsDifference(nums []int, goal int) int {
	out := Min(Abs(goal), Abs(nums[0]-goal))
	n := len(nums)
	if n == 1 {
		return out
	}

	m := n >> 1
	ls, rs := m, n-m
	lSum := make([]int, 1<<ls)
	rSum := make([]int, 1<<rs)
	for i := 1; i < 1<<ls; i++ {
		for j := 0; j < ls; j++ {
			if i&(1<<j) > 0 {
				lSum[i] = lSum[i-(1<<j)] + nums[j]
				break
			}
		}
	}
	for i := 1; i < 1<<rs; i++ {
		for j := 0; j < rs; j++ {
			if i&(1<<j) > 0 {
				rSum[i] = rSum[i-(1<<j)] + nums[ls+j]
				break
			}
		}
	}

	for _, sum := range lSum {
		out = Min(out, Abs(sum-goal))
	}
	for _, sum := range rSum {
		out = Min(out, Abs(sum-goal))
	}
	sort.Ints(lSum)
	sort.Ints(rSum)

	fmt.Println(lSum)
	fmt.Println(rSum)

	for i, j := 0, 1<<rs-1; i < 1<<ls && j >= 0; {
		sum := lSum[i] + rSum[j]
		out = Min(out, Abs(sum-goal))
		if sum <= goal {
			i++
		} else {
			j--
		}
	}
	return out
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
