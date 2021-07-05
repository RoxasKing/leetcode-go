package main

func numTimesAllBlue(light []int) int {
	out, max := 0, 0
	for i, idx := range light {
		max = Max(max, idx)
		if max == i+1 {
			out++
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
