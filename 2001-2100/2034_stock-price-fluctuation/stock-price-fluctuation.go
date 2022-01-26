package main

import "container/heap"

// Difficulty:
// Medium

// Tags:
// Hash
// Priority Queue

type StockPrice struct {
	time, cur int
	h         map[int]int
	max       MaxHeap
	min       MinHeap
}

func Constructor() StockPrice {
	return StockPrice{h: map[int]int{}}
}

func (this *StockPrice) Update(timestamp int, price int) {
	if this.time <= timestamp {
		this.time = timestamp
		this.cur = price
	}
	this.h[timestamp] = price
	for ; this.max.Len() > 0 && this.h[this.max[0][0]] != this.max[0][1]; heap.Pop(&this.max) {
	}
	for ; this.min.Len() > 0 && this.h[this.min[0][0]] != this.min[0][1]; heap.Pop(&this.min) {
	}
	heap.Push(&this.max, [2]int{timestamp, price})
	heap.Push(&this.min, [2]int{timestamp, price})
}

func (this *StockPrice) Current() int {
	return this.cur
}

func (this *StockPrice) Maximum() int {
	return this.max[0][1]
}

func (this *StockPrice) Minimum() int {
	return this.min[0][1]
}

type MinHeap [][2]int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i][1] < h[j][1] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *MinHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}

type MaxHeap [][2]int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i][1] > h[j][1] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *MaxHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
