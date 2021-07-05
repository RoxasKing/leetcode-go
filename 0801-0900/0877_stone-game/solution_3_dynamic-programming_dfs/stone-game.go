package main

// Tags:
// Dynamic Programming
// DFS

func stoneGame(piles []int) bool {
	l, r := 0, len(piles)-1
	return dp(piles, 0, 0, l, r)
}

func dp(piles []int, a, b, l, r int) bool {
	if l > r {
		return a > b
	}
	return dp(piles, a+piles[l], b, l+1, r) || dp(piles, a+piles[r], b, l, r-1)
}
