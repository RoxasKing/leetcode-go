package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func maxProfit(prices []int) int {
	sell, buy := 0, -1<<31
	for _, price := range prices {
		sell = Max(sell, buy+price)
		buy = Max(buy, sell-price)
	}
	return sell
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
