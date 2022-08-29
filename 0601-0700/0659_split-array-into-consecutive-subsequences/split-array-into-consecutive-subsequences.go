package main

import "container/heap"

// Difficulty:
// Medium

// Tags:
// Hash
// Priority Queue

func isPossible(nums []int) bool {
	hs := map[int]*minh{}
	for _, x := range nums {
		if hs[x] == nil {
			hs[x] = &minh{}
		}
		if h := hs[x-1]; h == nil {
			heap.Push(hs[x], 1)
		} else {
			v := heap.Pop(h).(int)
			if h.Len() == 0 {
				delete(hs, x-1)
			}
			heap.Push(hs[x], v+1)
		}
	}
	for _, h := range hs {
		if (*h)[0] < 3 {
			return false
		}
	}
	return true
}

type minh []int

func (h minh) Len() int            { return len(h) }
func (h minh) Less(i, j int) bool  { return h[i] < h[j] }
func (h minh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minh) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minh) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }
