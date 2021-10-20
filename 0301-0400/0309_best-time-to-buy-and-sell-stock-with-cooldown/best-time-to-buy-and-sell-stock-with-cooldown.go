package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func maxProfit(prices []int) int {
	preSell, sell, buy := 0, 0, -1<<31
	for _, p := range prices {
		tmp := sell
		sell = Max(sell, buy+p)
		buy = Max(buy, preSell-p)
		preSell = tmp
	}
	return sell
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
