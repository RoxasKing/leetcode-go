package main

// Difficulty:
// Easy

// Tags:
// Dynamic Programming

func maxProfit(prices []int) int {
	sell, buy := 0, -1<<31
	for _, p := range prices {
		sell = Max(sell, buy+p)
		buy = Max(buy, -p)
	}
	return sell
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
