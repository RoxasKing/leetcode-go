package main

import (
	"container/heap"
)

// Tags:
// Sliding Window + Monotone Stack
func longestSubarray(nums []int, limit int) int {
	out := 0
	mins := Stack{}
	maxs := Stack{}
	n := len(nums)

	for l, r := 0, 0; r < n; r++ {
		for mins.Len() > 0 && nums[r] < mins.Top() {
			mins.Pop()
		}
		mins.Push(nums[r])

		for maxs.Len() > 0 && nums[r] > maxs.Top() {
			maxs.Pop()
		}
		maxs.Push(nums[r])

		for l < r && maxs[0]-mins[0] > limit {
			if mins[0] == nums[l] {
				mins = mins[1:]
			}
			if maxs[0] == nums[l] {
				maxs = maxs[1:]
			}
			l++
		}

		out = Max(out, r+1-l)
	}
	return out
}

type Stack []int

func (s Stack) Len() int    { return len(s) }
func (s Stack) Top() int    { return s[len(s)-1] }
func (s *Stack) Push(x int) { *s = append(*s, x) }
func (s *Stack) Pop() int {
	last := len(*s) - 1
	out := (*s)[last]
	*s = (*s)[:last]
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Sliding Window + Priority Queue
func longestSubarray2(nums []int, limit int) int {
	out := 0
	n := len(nums)
	h := NewAbsHeap()
	for l, r := 0, 0; r < n; r++ {
		h.Push(nums[r])
		for l < r && h.GetAbs() > limit {
			h.Erase(nums[l])
			l++
		}
		out = Max(out, r+1-l)
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

type AbsHeap struct {
	maxh MaxHeap
	minh MinHeap
	maxd map[int]int
	mind map[int]int
}

func NewAbsHeap() *AbsHeap {
	return &AbsHeap{
		maxh: MaxHeap{},
		minh: MinHeap{},
		maxd: make(map[int]int),
		mind: make(map[int]int),
	}
}

func (h *AbsHeap) Push(x int) {
	heap.Push(&h.maxh, x)
	heap.Push(&h.minh, x)
}

func (h *AbsHeap) Erase(x int) {
	h.maxd[x]++
	h.mind[x]++
	for h.maxh.Len() > 0 && h.maxd[h.maxh[0]] > 0 {
		h.maxd[h.maxh[0]]--
		heap.Pop(&h.maxh)
	}
	for h.minh.Len() > 0 && h.mind[h.minh[0]] > 0 {
		h.mind[h.minh[0]]--
		heap.Pop(&h.minh)
	}
}

func (h *AbsHeap) GetAbs() int {
	return h.maxh[0] - h.minh[0]
}
