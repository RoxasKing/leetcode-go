package main

import "container/heap"

// Tags:
// Priority Queue

type MedianFinder struct {
	maxh MaxHeap
	minh MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxh.Len() == 0 || this.maxh[0] > num {
		heap.Push(&this.maxh, num)
	} else {
		heap.Push(&this.minh, num)
	}
	if this.maxh.Len() < this.minh.Len() {
		heap.Push(&this.maxh, heap.Pop(&this.minh))
	} else if this.maxh.Len() > this.minh.Len()+1 {
		heap.Push(&this.minh, heap.Pop(&this.maxh))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxh.Len() > this.minh.Len() {
		return float64(this.maxh[0])
	}
	return float64(this.maxh[0]+this.minh[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}
