package main

import "container/heap"

// Priority Queue(Heap Sort)

type SeatManager struct {
	h      *MinHeap
	inHeap map[int]bool
}

func Constructor(n int) SeatManager {
	h := &MinHeap{}
	inHeap := make(map[int]bool)
	for i := 1; i <= n; i++ {
		heap.Push(h, i)
		inHeap[i] = true
	}
	return SeatManager{
		h:      h,
		inHeap: inHeap,
	}
}

func (this *SeatManager) Reserve() int {
	num := heap.Pop(this.h).(int)
	this.inHeap[num] = false
	return num
}

func (this *SeatManager) Unreserve(seatNumber int) {
	if this.inHeap[seatNumber] {
		return
	}
	this.inHeap[seatNumber] = true
	heap.Push(this.h, seatNumber)
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	last := len(*h) - 1
	out := (*h)[last]
	*h = (*h)[:last]
	return out
}
