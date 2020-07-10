package leetcode

/*
  给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。
  你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
  返回获得利润的最大值。
  注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。

  注意:
    0 < prices.length <= 50000.
    0 < prices[i] < 50000.
    0 <= fee < 50000.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxProfitVI(prices []int, fee int) int {
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		dp0 = Max(dp0, dp1+prices[i]-fee)
		dp1 = Max(dp1, dp0-prices[i])
	}
	return dp0
}
