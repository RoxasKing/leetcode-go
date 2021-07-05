package main

func largestAltitude(gain []int) int {
	out, sum := 0, 0
	for _, g := range gain {
		sum += g
		out = Max(out, sum)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
