package main

import (
	"container/heap"
	"sort"
)

// Difficulty:
// Medium

// Tags:
// Priority Queue

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := MaxHeap{}
	for _, u := range nums1 {
		for _, v := range nums2 {
			if h.Len() == k {
				if sum := h[0][0] + h[0][1]; sum <= u+v {
					break
				} else if sum > u+v {
					heap.Pop(&h)
				}
			}
			heap.Push(&h, []int{u, v})
		}
	}
	sort.Slice(h, func(i, j int) bool {
		if sumi, sumj := h[i][0]+h[i][1], h[j][0]+h[j][1]; sumi != sumj {
			return sumi < sumj
		}
		return h[i][0] < h[j][0] || h[i][0] == h[j][0] && h[i][1] < h[j][i]
	})
	return h
}

type MaxHeap [][]int                  // [u, v] pair
func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i][0]+h[i][1] > h[j][0]+h[j][1] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *MaxHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}
