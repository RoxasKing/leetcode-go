package leetcode

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
func maxProfitVI(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][3]int, len(prices))
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][2]-prices[i]) // 持有股票
		dp[i][1] = dp[i-1][0] + prices[i]                // 不持有股票，冷冻期
		dp[i][2] = Max(dp[i-1][1], dp[i-1][2])           // 不持有股票，非冷冻期
	}
	return Max(dp[len(prices)-1][1], dp[len(prices)-1][2])
}

// Dynamic Programming
func maxProfitVI2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := [2][3]int{{-prices[0]}}
	for i := 1; i < len(prices); i++ {
		dp[1][0] = Max(dp[0][0], dp[0][2]-prices[i])
		dp[1][1] = dp[0][0] + prices[i]
		dp[1][2] = Max(dp[0][1], dp[0][2])
		dp[0] = dp[1]
	}
	return Max(dp[1][1], dp[1][2])
}
