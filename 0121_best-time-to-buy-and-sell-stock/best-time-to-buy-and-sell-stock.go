package main

/*
  给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
  如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
  注意：你不能在买入股票前卖出股票。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxProfit(prices []int) int {
	profit, buyPrice := 0, -1<<31
	for _, price := range prices {
		profit = Max(profit, buyPrice+price)
		buyPrice = Max(buyPrice, -price)
	}
	return profit
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
