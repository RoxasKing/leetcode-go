package main

// Difficulty:
// Medium

// Tags:
// Greedy

func minDominoRotations(tops []int, bottoms []int) int {
	n := len(tops)
	freq := [7]int{}
	for i := 0; i < n; i++ {
		freq[tops[i]]++
		if tops[i] != bottoms[i] {
			freq[bottoms[i]]++
		}
	}
	out := -1
	for k := 1; k <= 6; k++ {
		if freq[k] != n {
			continue
		}
		c1, c2 := 0, 0
		for i := 0; i < n; i++ {
			if tops[i] != k {
				c1++
			}
			if bottoms[i] != k {
				c2++
			}
		}
		if out == -1 || out > c1 {
			out = c1
		}
		if out > c2 {
			out = c2
		}
	}
	return out
}
