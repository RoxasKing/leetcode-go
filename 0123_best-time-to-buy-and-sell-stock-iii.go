package leetcode

/*
  给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
  设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
  注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxProfitIII(prices []int) int {
	dp := [3][2]int{{0, 0}, {0, -1 << 31}, {0, -1 << 31}}
	for _, price := range prices {
		for i := 2; i >= 1; i-- {
			dp[i][0] = Max(dp[i][0], dp[i][1]+price)
			dp[i][1] = Max(dp[i][1], dp[i-1][0]-price)
		}
	}
	return dp[2][0]
}
