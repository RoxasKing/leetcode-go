package main

/*
  编写一个 StockSpanner 类，它收集某些股票的每日报价，并返回该股票当日价格的跨度。
  今天股票价格的跨度被定义为股票价格小于或等于今天价格的最大连续日数（从今天开始往回数，包括今天）。
  例如，如果未来7天股票的价格是 [100, 80, 60, 70, 60, 75, 85]，那么股票跨度将是 [1, 1, 1, 2, 1, 4, 6]。

  提示：
    调用 StockSpanner.next(int price) 时，将有 1 <= price <= 10^5。
    每个测试用例最多可以调用  10000 次 StockSpanner.next。
    在所有测试用例中，最多调用 150000 次 StockSpanner.next。
    此问题的总时间限制减少了 50%。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/online-stock-span
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Monotone Stack
type StockSpanner struct {
	prices []int
	counts []int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	count := 1
	for len(this.prices) > 0 && this.prices[len(this.prices)-1] <= price {
		count += this.counts[len(this.counts)-1]
		this.counts = this.counts[:len(this.counts)-1]
		this.prices = this.prices[:len(this.prices)-1]
	}
	this.prices = append(this.prices, price)
	this.counts = append(this.counts, count)
	return count
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := NewStockSpanner();
 * param_1 := obj.Next(price);
 */
