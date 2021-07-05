package main

import "container/heap"

// Priority Queue(Heap Sort)

type MedianFinder struct {
	maxh MaxHeap
	minh MinHeap
	size int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		maxh: MaxHeap{},
		minh: MinHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.maxh, num)
	this.size++
	for {
		if this.maxh.Len() > 0 && this.minh.Len() > 0 && this.maxh[0] > this.minh[0] {
			a := heap.Pop(&this.maxh)
			b := heap.Pop(&this.minh)
			heap.Push(&this.maxh, b)
			heap.Push(&this.minh, a)
		} else if this.maxh.Len() > this.minh.Len()+1 {
			heap.Push(&this.minh, heap.Pop(&this.maxh))
		} else {
			break
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.size&1 == 1 {
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
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
