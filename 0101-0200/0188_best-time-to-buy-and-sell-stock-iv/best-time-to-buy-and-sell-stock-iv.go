package main

/*
  You are given an integer array prices where prices[i] is the price of a given stock on the ith day, and an integer k.

  Find the maximum profit you can achieve. You may complete at most k transactions.

  Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

  Example 1:
    Input: k = 2, prices = [2,4,1]
    Output: 2
    Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.

  Example 2:
    Input: k = 2, prices = [3,2,6,5,0,3]
    Output: 7
    Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4. Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.

  Constraints:
    1. 0 <= k <= 100
    2. 0 <= prices.length <= 1000
    3. 0 <= prices[i] <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

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
