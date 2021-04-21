package main

/*
  给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
  设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
  注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

  示例 1：
    输入：k = 2, prices = [2,4,1]
    输出：2
    解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。

  示例 2：
    输入：k = 2, prices = [3,2,6,5,0,3]
    输出：7
    解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
         随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。

  提示：
    0 <= k <= 109
    0 <= prices.length <= 1000
    0 <= prices[i] <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) < 2 {
		return 0
	} else if k >= len(prices)>>1 { // 如果交易次数大于最多交易次数(天数的一半)，可以看作是没有交易限制
		return maxProfitWithNoLimit(prices)
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

func maxProfitWithNoLimit(prices []int) int {
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
