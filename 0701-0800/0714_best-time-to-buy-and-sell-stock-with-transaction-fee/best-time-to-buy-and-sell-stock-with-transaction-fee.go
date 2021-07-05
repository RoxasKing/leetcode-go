package main

// Tags:
// Dynamic Programming
func maxProfit(prices []int, fee int) int {
	dp0, dp1 := 0, -1<<31
	for _, price := range prices {
		dp0 = Max(dp0, dp1+price-fee)
		dp1 = Max(dp1, dp0-price)
	}
	return dp0
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
