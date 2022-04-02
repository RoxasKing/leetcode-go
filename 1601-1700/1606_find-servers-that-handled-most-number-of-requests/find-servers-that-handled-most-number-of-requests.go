package main

import "container/heap"

// Difficulty:
// Hard

// Tags:
// Priority Queue

func busiestServers(k int, arrival []int, load []int) []int {
	a, b := make(availableH, k), make(busyH, 0, k)
	for i := 0; i < k; i++ {
		a[i] = i
	}
	cnt := make([]int, k)
	out, max := []int{}, 0
	for i, t := range arrival {
		for ; b.Len() > 0 && b[0].it <= t; heap.Push(&a, i+((heap.Pop(&b).(server).id-i)%k+k)%k) {
		}
		if a.Len() == 0 {
			continue
		}
		id := heap.Pop(&a).(int) % k
		if cnt[id]++; max < cnt[id] {
			out, max = []int{id}, cnt[id]
		} else if max == cnt[id] {
			out = append(out, id)
		}
		heap.Push(&b, server{t + load[i], id})
	}
	return out
}

type availableH []int

func (h availableH) Len() int            { return len(h) }
func (h availableH) Less(i, j int) bool  { return h[i] < h[j] }
func (h availableH) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *availableH) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *availableH) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }

type server struct{ it, id int }
type busyH []server

func (h busyH) Len() int            { return len(h) }
func (h busyH) Less(i, j int) bool  { return h[i].it < h[j].it }
func (h busyH) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *busyH) Push(x interface{}) { *h = append(*h, x.(server)) }
func (h *busyH) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }
