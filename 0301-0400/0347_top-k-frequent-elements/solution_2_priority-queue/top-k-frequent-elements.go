package main

import "container/heap"

// Difficulty:
// Medium

// Tags:
// Priority Queue

func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}
	h := minh{}
	for num, cnt := range freq {
		if h.Len() < k {
			heap.Push(&h, pair{num, cnt})
		} else if h.Len() == k && h[0].v < cnt {
			heap.Pop(&h)
			heap.Push(&h, pair{num, cnt})
		}
	}
	o := make([]int, k)
	for i := range o {
		o[i] = h[i].k
	}
	return o
}

type pair struct{ k, v int }
type minh []pair

func (h minh) Len() int            { return len(h) }
func (h minh) Less(i, j int) bool  { return h[i].v < h[j].v }
func (h minh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minh) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *minh) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }
