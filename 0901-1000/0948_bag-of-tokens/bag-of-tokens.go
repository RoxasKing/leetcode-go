package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Greedy
// Sorting
// Two Pointers

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	o := 0
	for l, r, score := 0, len(tokens)-1, 0; l <= r; {
		if power >= tokens[l] {
			power -= tokens[l]
			l++
			score++
			o = max(o, score)
		} else if score > 0 {
			power += tokens[r]
			r--
			score--
		} else {
			break
		}
	}
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
