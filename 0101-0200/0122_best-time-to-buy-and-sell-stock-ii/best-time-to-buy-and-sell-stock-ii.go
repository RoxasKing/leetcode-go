package main

// Tags:
// Dynamic Programming
func maxProfit(prices []int) int {
	sell, buy := 0, -1<<31
	for _, price := range prices {
		preSell := sell
		sell = Max(sell, buy+price)
		buy = Max(buy, preSell-price)
	}
	return sell
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
