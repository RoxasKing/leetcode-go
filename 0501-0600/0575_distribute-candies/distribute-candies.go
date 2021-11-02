package main

// Difficulty:
// Easy

// Tags:
// Hash

func distributeCandies(candyType []int) int {
	n := len(candyType) >> 1
	vis := map[int]struct{}{}
	out := 0
	for _, t := range candyType {
		if _, ok := vis[t]; !ok {
			vis[t] = struct{}{}
			out++
		}
	}
	if out > n {
		out = n
	}
	return out
}
