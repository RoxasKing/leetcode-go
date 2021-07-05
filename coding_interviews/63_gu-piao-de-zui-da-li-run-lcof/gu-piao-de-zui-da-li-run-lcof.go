package main

// Tags:
// Dynamic Programming
func maxProfit(prices []int) int {
	out, cost := 0, -1<<31
	for _, p := range prices {
		out = Max(out, p+cost)
		cost = Max(cost, -p)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
