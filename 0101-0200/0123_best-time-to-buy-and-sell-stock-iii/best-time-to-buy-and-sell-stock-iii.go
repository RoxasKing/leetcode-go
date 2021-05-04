package main

/*
  You are given an array prices where prices[i] is the price of a given stock on the ith day.

  Find the maximum profit you can achieve. You may complete at most two transactions.

  Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

  Example 1:
    Input: prices = [3,3,5,0,0,3,1,4]
    Output: 6
    Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
    Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.

  Example 2:
    Input: prices = [1,2,3,4,5]
    Output: 4
    Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
    Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at the same time. You must sell before buying again.

  Example 3:
    Input: prices = [7,6,4,3,1]
    Output: 0
    Explanation: In this case, no transaction is done, i.e. max profit = 0.

  Example 4:
    Input: prices = [1]
    Output: 0

  Constraints:
    1. 1 <= prices.length <= 10^5
    2. 0 <= prices[i] <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming

func maxProfit(prices []int) int {
	buy1, sell1, buy2, sell2 := -1<<31, 0, -1<<31, 0
	for _, price := range prices {
		buy1 = Max(buy1, -price)
		sell1 = Max(sell1, buy1+price)
		buy2 = Max(buy2, sell1-price)
		sell2 = Max(sell2, buy2+price)
	}
	return sell2
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
