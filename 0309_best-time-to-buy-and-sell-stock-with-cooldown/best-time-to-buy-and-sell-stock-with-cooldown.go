package main

/*
  给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​

  设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
    你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
    卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := [2][3]int{{0, 0, -1 << 31}}
	for i := 0; i < len(prices); i++ {
		dp[1][0] = Max(dp[0][0], dp[0][1])           // 不持有股票，非冷冻期
		dp[1][1] = dp[0][2] + prices[i]              // 不持有股票，冷冻期
		dp[1][2] = Max(dp[0][2], dp[0][0]-prices[i]) // 持有股票
		dp[0] = dp[1]
	}
	return Max(dp[1][0], dp[1][1])
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
