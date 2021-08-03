package main

// Tags:
// Greedy

func numberOfWeeks(milestones []int) int64 {
	sum, max := 0, 0
	for _, m := range milestones {
		sum += m
		max = Max(max, m)
	}
	if max <= sum>>1 {
		return int64(sum)
	}
	return int64((sum-max)<<1 + 1)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
