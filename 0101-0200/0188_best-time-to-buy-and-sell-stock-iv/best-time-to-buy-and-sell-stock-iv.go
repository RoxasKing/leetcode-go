package main

// Tags:
// Dynamic Programming
func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) < 2 {
		return 0
	} else if k >= len(prices)>>1 { // 如果交易次数大于最多交易次数(天数的一半)，可以看作是没有交易限制
		return maxProfitInf(prices)
	}
	dp := make([][2]int, k+1)
	for i := range dp {
		dp[i][0] = -1 << 31
	}
	for _, price := range prices {
		for i := 1; i <= k; i++ {
			dp[i][0] = Max(dp[i][0], dp[i-1][1]-price) // bul
			dp[i][1] = Max(dp[i][1], dp[i][0]+price)   // sell
		}
	}
	return dp[k][1] // last sell
}

func maxProfitInf(prices []int) int {
	buy, sell := -1<<31, 0
	for _, price := range prices {
		preBuy := buy
		buy = Max(buy, sell-price)
		sell = Max(sell, preBuy+price)
	}
	return sell
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
