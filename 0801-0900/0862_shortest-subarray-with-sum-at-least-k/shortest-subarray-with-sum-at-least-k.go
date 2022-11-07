package main

// Difficulty:
// Hard

// Tags:
// Prefix Sum
// Sliding Window
// Stack

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	ps := make([]int, n+1)
	for i := 0; i < n; i++ {
		ps[i+1] = ps[i] + nums[i]
	}
	o, q := -1, []int{}
	for i, x := range ps {
		for ; len(q) > 0 && x <= ps[q[len(q)-1]]; q = q[:len(q)-1] {
		}
		for ; len(q) > 0 && x-ps[q[0]] >= k; q = q[1:] {
			if o == -1 || o > i-q[0] {
				o = i - q[0]
			}
		}
		q = append(q, i)
	}
	return o
}
