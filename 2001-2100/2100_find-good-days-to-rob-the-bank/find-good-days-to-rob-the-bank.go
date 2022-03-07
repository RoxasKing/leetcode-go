package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	lr := make([]int, n)
	rl := make([]int, n)
	for i := 1; i < n; i++ {
		if security[i-1] >= security[i] {
			lr[i] = lr[i-1] + 1
		}
		if security[n-1-i] <= security[n-i] {
			rl[n-1-i] = rl[n-i] + 1
		}
	}
	out := []int{}
	for i := 0; i < n; i++ {
		if lr[i] >= time && rl[i] >= time {
			out = append(out, i)
		}
	}
	return out
}
