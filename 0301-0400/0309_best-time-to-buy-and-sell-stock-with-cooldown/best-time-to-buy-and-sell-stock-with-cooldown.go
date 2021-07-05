package main

// Tags:
// Dynamic Programming
func maxProfit(prices []int) int {
	dp0, dp1, dp2 := -1<<31, 0, 0 // dp0:买入 dp1:卖出,冷冻期 dp2:卖出,非冷冻期
	for i := 0; i < len(prices); i++ {
		dp0, dp1, dp2 = Max(dp0, dp2-prices[i]), dp0+prices[i], Max(dp1, dp2)
	}
	return Max(dp1, dp2)
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Dynamic Programming
func maxProfit2(prices []int) int {
	dp0, dp1, dpp0 := 0, -1<<31, 0
	for i := range prices {
		tmp := dp0
		dp0 = Max(dp0, dp1+prices[i])
		dp1 = Max(dp1, dpp0-prices[i])
		dpp0 = tmp
	}
	return dp0
}
