package main

import (
	"container/heap"
)

// Tags:
// Priority Queue

func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}

	h := &MinHeap{}
	for k, v := range cnt {
		heap.Push(h, &pair{num: k, cnt: v})
	}

	out := make([]int, 0, k)
	for len(out) < k {
		out = append(out, heap.Pop(h).(*pair).num)
	}
	return out
}

type pair struct {
	num, cnt int
}

type MinHeap []*pair

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].cnt > h[j].cnt }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*pair)) }
func (h *MinHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
