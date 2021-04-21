package main

/*
  给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
  设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

  注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

  提示：
    1 <= prices.length <= 3 * 10 ^ 4
    0 <= prices[i] <= 10 ^ 4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxProfit(prices []int) int {
	var profit int
	for i := 1; i < len(prices); i++ {
		profit += Max(0, prices[i]-prices[i-1])
	}
	return profit
}

func maxProfit2(prices []int) int {
	sel, buy := 0, -1<<31
	for _, price := range prices {
		preSel := sel
		sel = Max(sel, buy+price)
		buy = Max(buy, preSel-price)
	}
	return sel
}

/*
	sel > buy+price --> sel-price > buy --> buy < sel-price
		because preSell == sel --> preSel-price == sel-price
	sel < buy+price --> sel-price < buy --> buy > sel-price
		because preSel < sel --> sel-price > preSel-price --> buy > preSel-price
*/

func maxProfit3(prices []int) int {
	sel, buy := 0, -1<<31
	for _, price := range prices {
		sel = Max(sel, buy+price)
		buy = Max(buy, sel-price)
	}
	return sel
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
