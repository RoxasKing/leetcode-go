package main

// Difficulty:
// Easy

func maximumWealth(accounts [][]int) int {
	out := 0
	for i := range accounts {
		sum := 0
		for j := range accounts[i] {
			sum += accounts[i][j]
		}
		if out < sum {
			out = sum
		}
	}
	return out
}
