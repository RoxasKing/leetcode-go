package main

// Tags:
// Brain Teaser

func getLastMoment(n int, left []int, right []int) int {
	lMax, rMax := 0, 0
	for _, idx := range left {
		lMax = Max(lMax, idx)
	}
	for _, idx := range right {
		rMax = Max(rMax, n-idx)
	}
	return Max(lMax, rMax)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
