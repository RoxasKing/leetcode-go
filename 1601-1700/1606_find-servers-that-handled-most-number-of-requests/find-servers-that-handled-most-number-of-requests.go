package main

import (
	"container/heap"
	"sort"
)

// Difficulty:
// Hard

// Tags:
// Priority Queue

func busiestServers(k int, arrival []int, load []int) []int {
	a, b := hi{make([]int, k)}, hs{} // available servers' idx ; busy servers' state
	for i := 0; i < k; i++ {
		a.IntSlice[i] = i
	}
	max := 0
	cnt := make([]int, k)
	out := []int{}
	for i, t := range arrival {
		for len(b) > 0 && b[0].idl <= t {
			heap.Push(&a, i+((b[0].idx-i)%k+k)%k) // Attention!!
			heap.Pop(&b)
		}
		if a.Len() == 0 {
			continue
		}
		idx := heap.Pop(&a).(int) % k
		if cnt[idx]++; max == cnt[idx] {
			out = append(out, idx)
		} else if max < cnt[idx] {
			max = cnt[idx]
			out = []int{idx}
		}
		heap.Push(&b, server{t + load[i], idx})
	}
	return out
}

type hi struct{ sort.IntSlice }

func (h *hi) Push(x interface{}) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *hi) Pop() interface{} {
	a := h.IntSlice
	i := a.Len() - 1
	o := a[i]
	h.IntSlice = a[:i]
	return o
}

type server struct{ idl, idx int } // server's idle time and idx
type hs []server

func (h hs) Len() int            { return len(h) }
func (h hs) Less(i, j int) bool  { return h[i].idl < h[j].idl }
func (h hs) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hs) Push(x interface{}) { *h = append(*h, x.(server)) }
func (h *hs) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }
