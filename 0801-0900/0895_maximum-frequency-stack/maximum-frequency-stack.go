package main

import "container/heap"

// Priority Queue

type FreqStack struct {
	cnt map[int]int
	idx int
	h   MaxHeap
}

func Constructor() FreqStack {
	return FreqStack{
		cnt: map[int]int{},
		idx: -1,
		h:   MaxHeap{},
	}
}

func (this *FreqStack) Push(val int) {
	this.cnt[val]++
	this.idx++
	heap.Push(&this.h, [3]int{val, this.cnt[val], this.idx})
}

func (this *FreqStack) Pop() int {
	top := heap.Pop(&this.h).([3]int)
	this.cnt[top[0]]--
	return top[0]
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */

type MaxHeap [][3]int // [val, freq, lastIdx]

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i][1] != h[j][1] {
		return h[i][1] > h[j][1]
	}
	return h[i][2] > h[j][2]
}
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.([3]int)) }
func (h *MaxHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
