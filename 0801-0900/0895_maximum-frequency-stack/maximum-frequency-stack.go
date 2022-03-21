package main

import "container/heap"

// Difficulty:
// Hard

// Tags:
// Priority Queue

type FreqStack struct {
	c map[int]int
	h MaxHeap
	i int
}

func Constructor() FreqStack {
	return FreqStack{c: map[int]int{}}
}

func (this *FreqStack) Push(val int) {
	this.c[val]++
	this.i++
	heap.Push(&this.h, Entry{val, this.c[val], this.i})
}

func (this *FreqStack) Pop() int {
	e := heap.Pop(&this.h).(Entry)
	this.c[e.v]--
	return e.v
}

type Entry struct{ v, c, i int }
type MaxHeap []Entry

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].c > h[j].c || h[i].c == h[j].c && h[i].i > h[j].i }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Entry)) }
func (h *MaxHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
