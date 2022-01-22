package main

// Difficulty:
// Medium

// Tags:
// Math
// Game Theory

func stoneGameIX(stones []int) bool {
	freq := [3]int{}
	for _, stone := range stones {
		freq[stone%3]++
	}
	c0, c1, c2 := freq[0], freq[1], freq[2]
	return c0&1 == 0 && c1 >= 1 && c2 >= 1 || c0&1 == 1 && (c1-c2 > 2 || c2-c1 > 2)
}
