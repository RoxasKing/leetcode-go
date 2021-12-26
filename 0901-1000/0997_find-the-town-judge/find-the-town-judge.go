package main

// Difficulty:
// Easy

// Tags:
// Hash

func findJudge(n int, trust [][]int) int {
	freq := make([]int, n+1)
	for _, t := range trust {
		a, b := t[0], t[1]
		freq[a]--
		freq[b]++
	}
	for i := 1; i <= n; i++ {
		if freq[i] == n-1 {
			return i
		}
	}
	return -1
}
