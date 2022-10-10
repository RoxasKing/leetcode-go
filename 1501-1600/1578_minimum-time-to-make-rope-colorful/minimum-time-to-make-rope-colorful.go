package main

// Difficulty:
// Medium

// Tags:
// Greedy

func minCost(colors string, neededTime []int) int {
	o, t := neededTime[0], neededTime[0]
	for i := 1; i < len(colors); i++ {
		o += neededTime[i]
		if colors[i-1] == colors[i] {
			t = max(t, neededTime[i])
			continue
		}
		o -= t
		t = neededTime[i]
	}
	o -= t
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
