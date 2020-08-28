package main

/*
  给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
  设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
  注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxProfitIV(k int, prices []int) int {
	if k == 0 || len(prices) < 2 {
		return 0
	} else if k >= len(prices)>>1 {
		return maxProfitII(prices)
	}
	dp := make([][2]int, k+1)
	for i := range dp {
		dp[i][1] = -1 << 31
	}
	for _, price := range prices {
		for i := k; i >= 1; i-- {
			dp[i][0] = Max(dp[i][0], dp[i][1]+price)
			dp[i][1] = Max(dp[i][1], dp[i-1][0]-price)
		}
	}
	return dp[k][0]
}
