package main

// Difficulty:
// Hard

// Tags:
// Backtracking

func maximumRequests(n int, requests [][]int) int {
	out, cnt, ful := 0, 0, n
	h := make([]int, n)
	var backtrack func(int)
	backtrack = func(i int) {
		if i == len(requests) {
			if ful == n {
				out = Max(out, cnt)
			}
			return
		}
		dec, inc := requests[i][0], requests[i][1]
		if h[dec]--; h[dec] == 0 {
			ful++
		} else if h[dec] == -1 {
			ful--
		}
		if h[inc]++; h[inc] == 0 {
			ful++
		} else if h[inc] == 1 {
			ful--
		}
		cnt++
		backtrack(i + 1)
		cnt--
		if h[dec]++; h[dec] == 0 {
			ful++
		} else if h[dec] == 1 {
			ful--
		}
		if h[inc]--; h[inc] == 0 {
			ful++
		} else if h[inc] == -1 {
			ful--
		}
		backtrack(i + 1)
	}
	backtrack(0)
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
