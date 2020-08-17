package leetcode

/*
  给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
  设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
  注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/

func maxProfitII(prices []int) int {
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

func maxProfitII2(prices []int) int {
	sel, buy := 0, -1<<31
	for _, price := range prices {
		sel = Max(sel, buy+price)
		buy = Max(buy, sel-price)
	}
	return sel
}
