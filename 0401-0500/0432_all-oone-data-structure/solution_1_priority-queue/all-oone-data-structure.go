package main

import "container/heap"

// Priority Queue

type AllOne struct {
	h    map[string]int
	maxh MaxHeap
	minh MinHeap
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		h:    map[string]int{},
		maxh: MaxHeap{},
		minh: MinHeap{},
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	this.h[key]++
	for this.maxh.Len() > 0 && this.maxh[0].cnt != this.h[this.maxh[0].key] {
		heap.Pop(&this.maxh)
	}
	for this.minh.Len() > 0 && this.minh[0].cnt != this.h[this.minh[0].key] {
		heap.Pop(&this.minh)
	}
	heap.Push(&this.maxh, &Pair{key: key, cnt: this.h[key]})
	heap.Push(&this.minh, &Pair{key: key, cnt: this.h[key]})
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	this.h[key]--
	for this.maxh.Len() > 0 && this.maxh[0].cnt != this.h[this.maxh[0].key] {
		heap.Pop(&this.maxh)
	}
	for this.minh.Len() > 0 && this.minh[0].cnt != this.h[this.minh[0].key] {
		heap.Pop(&this.minh)
	}
	if this.h[key] == 0 {
		delete(this.h, key)
		return
	}
	heap.Push(&this.maxh, &Pair{key: key, cnt: this.h[key]})
	heap.Push(&this.minh, &Pair{key: key, cnt: this.h[key]})
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.maxh.Len() > 0 {
		return this.maxh[0].key
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.minh.Len() > 0 {
		return this.minh[0].key
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */

type Pair struct {
	key string
	cnt int
}

type MaxHeap []*Pair

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].cnt > h[j].cnt }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(*Pair)) }
func (h *MaxHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}

type MinHeap []*Pair

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].cnt < h[j].cnt }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*Pair)) }
func (h *MinHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
