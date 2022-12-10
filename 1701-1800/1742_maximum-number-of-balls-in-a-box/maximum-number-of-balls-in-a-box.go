package main

// Difficulty:
// Easy

// Tags:
// Hash

func countBalls(lowLimit int, highLimit int) int {
	h := [46]int{}
	for i := lowLimit; i <= highLimit; i++ {
		x := 0
		for j := i; j > 0; j /= 10 {
			x += j % 10
		}
		h[x]++
	}
	o := 0
	for i := 1; i < 46; i++ {
		o = max(o, h[i])
	}
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
