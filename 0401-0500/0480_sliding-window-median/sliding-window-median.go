package main

import (
	"container/heap"
)

// Tags:
// Priority Queue(Heap Sort) + Sliding Window
func medianSlidingWindow(nums []int, k int) []float64 {
	n := len(nums)
	out := make([]float64, 0, n+1-k)
	midh := NewMidHeap()
	for i := range nums {
		midh.Push(nums[i])
		if i > k-1 {
			midh.Erase(nums[i-k])
		}
		if i >= k-1 {
			out = append(out, midh.GetMid())
		}
	}
	return out
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	last := len(*h) - 1
	out := (*h)[last]
	*h = (*h)[:last]
	return out
}

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

type MidHeap struct {
	mark map[int]int // count the number of times the number needs to be deleted
	maxh MaxHeap
	minh MinHeap
	maxd int // count the number of times needs to be deleted in max heap
	mind int // count the number of times needs to be deleted in min heap
}

func NewMidHeap() *MidHeap {
	return &MidHeap{
		mark: make(map[int]int),
		maxh: MaxHeap{},
		minh: MinHeap{},
	}
}

func (h *MidHeap) Push(x int) {
	if h.maxh.Len() == 0 || x <= h.maxh[0] {
		heap.Push(&h.maxh, x)
	} else {
		heap.Push(&h.minh, x)
	}
	h.balance()
}

func (h *MidHeap) Erase(x int) {
	h.mark[x]++
	if x <= h.maxh[0] {
		h.maxd++
	} else {
		h.mind++
	}
	h.balance()
}

func (h *MidHeap) balance() {
	for {
		for h.maxh.Len() > 0 && h.mark[h.maxh[0]] > 0 {
			h.mark[h.maxh[0]]--
			h.maxd--
			heap.Pop(&h.maxh)
		}
		for h.minh.Len() > 0 && h.mark[h.minh[0]] > 0 {
			h.mark[h.minh[0]]--
			h.mind--
			heap.Pop(&h.minh)
		}
		if h.maxh.Len()-h.maxd < h.minh.Len()-h.mind {
			heap.Push(&h.maxh, heap.Pop(&h.minh))
		} else if h.maxh.Len()-h.maxd > h.minh.Len()-h.mind+1 {
			heap.Push(&h.minh, heap.Pop(&h.maxh))
		} else {
			break
		}
	}
}

func (h *MidHeap) GetMid() float64 {
	if (h.maxh.Len()-h.maxd+h.minh.Len()-h.mind)&1 == 1 {
		return float64(h.maxh[0])
	}
	return float64(h.maxh[0]+h.minh[0]) / 2.0
}
