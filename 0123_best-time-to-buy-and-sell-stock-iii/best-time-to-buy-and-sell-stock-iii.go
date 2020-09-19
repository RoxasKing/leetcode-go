package main

/*
  给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
  设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
  注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func maxProfit(prices []int) int {
	dp := [3][2]int{{0, 0}, {0, -1 << 31}, {0, -1 << 31}}
	for _, price := range prices {
		for i := 2; i >= 1; i-- {
			dp[i][0] = Max(dp[i][0], dp[i][1]+price)   // sell
			dp[i][1] = Max(dp[i][1], dp[i-1][0]-price) // buy
		}
	}
	return dp[2][0]
}

// Backtracking
func maxProfit2(prices []int) int {
	memo := make([][][]int, len(prices)+1)
	for i := range memo {
		memo[i] = make([][]int, 2)
		for j := range memo[i] {
			memo[i][j] = []int{-1 << 31, -1 << 31, -1 << 31}
		}
	}
	return backtrack(prices, 0, 0, 0, memo)
}

func backtrack(prices []int, index, status, k int, memo [][][]int) int {
	if memo[index][status][k] != -1<<31 {
		return memo[index][status][k]
	}
	if k == 2 || index == len(prices) {
		return 0
	}

	var a, b, c int
	a = backtrack(prices, index+1, status, k, memo) // 不做任何处理
	if status == 0 {
		b = backtrack(prices, index+1, 1, k, memo) - prices[index] // 当前状态为卖出时，选择买入
	} else {
		c = backtrack(prices, index+1, 0, k+1, memo) + prices[index] // 当前状态为买入时，选择卖出
	}

	memo[index][status][k] = Max(Max(a, b), c)
	return memo[index][status][k]
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
