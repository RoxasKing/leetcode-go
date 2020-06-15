package codinginterviews

/*
  假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？

  限制：
    0 <= 数组长度 <= 10^5
*/

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	min, out := 1<<31-1, -1<<31
	for _, price := range prices {
		out = Max(out, price-min)
		min = Min(min, price)
	}
	if out < 0 {
		return 0
	}
	return out
}
