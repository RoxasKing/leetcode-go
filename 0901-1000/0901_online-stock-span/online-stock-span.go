package main

// Tags:
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
