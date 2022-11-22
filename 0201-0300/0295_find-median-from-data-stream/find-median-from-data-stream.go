package main

import "container/heap"

// Difficulty:
// Hard

// Tags:
// Priority Queue

type MedianFinder struct {
	max maxh
	min minh
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if this.max.Len() == 0 || this.max[0] > num {
		heap.Push(&this.max, num)
	} else {
		heap.Push(&this.min, num)
	}
	if this.max.Len() < this.min.Len() {
		heap.Push(&this.max, heap.Pop(&this.min))
	} else if this.max.Len() > this.min.Len()+1 {
		heap.Push(&this.min, heap.Pop(&this.max))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.max.Len() > this.min.Len() {
		return float64(this.max[0])
	}
	return float64(this.max[0]+this.min[0]) / 2
}

type maxh []int

func (h maxh) Len() int            { return len(h) }
func (h maxh) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxh) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxh) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }

type minh []int

func (h minh) Len() int            { return len(h) }
func (h minh) Less(i, j int) bool  { return h[i] < h[j] }
func (h minh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minh) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minh) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
