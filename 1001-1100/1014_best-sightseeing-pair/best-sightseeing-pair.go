package main

func maxScoreSightseeingPair(values []int) int {
	n := len(values)
	out := 0
	max := 0
	for i := 0; i < n; i++ {
		out = Max(out, max+values[i]-i) // (values[i] + i) + (values[j] - j)
		max = Max(max, values[i]+i)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
