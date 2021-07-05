package main

// Tags:
// Dynamic Programming

func maxProfit(prices []int) int {
	buy1, sell1, buy2, sell2 := -1<<31, 0, -1<<31, 0
	for _, price := range prices {
		buy1 = Max(buy1, -price)
		sell1 = Max(sell1, buy1+price)
		buy2 = Max(buy2, sell1-price)
		sell2 = Max(sell2, buy2+price)
	}
	return sell2
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
