package main

import (
	"math/rand"
)

// Difficulty:
// Medium

// Tags:
// 3-Way Quick Sort

func wiggleSort(nums []int) {
	n := len(nums)
	k := (n - 1) >> 1
	for l, r := 0, n-1; l < r; {
		pi := rand.Intn(r+1-l) + l
		pv := nums[pi]
		s, t := l, r
		for i := s; i <= t; {
			if nums[i] < pv {
				nums[i], nums[s] = nums[s], nums[i]
				i++
				s++
			} else if nums[i] > pv {
				nums[i], nums[t] = nums[t], nums[i]
				t--
			} else {
				i++
			}
		}
		if t < k {
			l = t + 1
		} else {
			if nums[t] == nums[k] {
				break
			}
			r = t
		}
	}
	for l, r := 0, k; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
	for l, r := k+1, n-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
	tmp := make([]int, n)
	copy(tmp, nums)
	for i := 0; i < n; i += 2 {
		nums[i] = tmp[0]
		tmp = tmp[1:]
	}
	for i := 1; i < n; i += 2 {
		nums[i] = tmp[0]
		tmp = tmp[1:]
	}
}
