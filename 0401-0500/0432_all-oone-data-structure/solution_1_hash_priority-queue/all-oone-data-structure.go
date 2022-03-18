package main

import "container/heap"

// Difficulty:
// Hard

// Tags:
// Hash
// Priority Queue

type AllOne struct {
	freq map[string]int
	maxh MaxHeap
	minh MinHeap
}

func Constructor() AllOne {
	return AllOne{freq: map[string]int{}}
}

func (this *AllOne) Inc(key string) {
	this.freq[key]++
	for ; this.maxh.Len() > 0 && this.maxh[0].v != this.freq[this.maxh[0].k]; heap.Pop(&this.maxh) {
	}
	for ; this.minh.Len() > 0 && this.minh[0].v != this.freq[this.minh[0].k]; heap.Pop(&this.minh) {
	}
	heap.Push(&this.maxh, &Pair{key, this.freq[key]})
	heap.Push(&this.minh, &Pair{key, this.freq[key]})
}

func (this *AllOne) Dec(key string) {
	this.freq[key]--
	for ; this.maxh.Len() > 0 && this.maxh[0].v != this.freq[this.maxh[0].k]; heap.Pop(&this.maxh) {
	}
	for ; this.minh.Len() > 0 && this.minh[0].v != this.freq[this.minh[0].k]; heap.Pop(&this.minh) {
	}
	if this.freq[key] == 0 {
		delete(this.freq, key)
		return
	}
	heap.Push(&this.maxh, &Pair{key, this.freq[key]})
	heap.Push(&this.minh, &Pair{key, this.freq[key]})
}

func (this *AllOne) GetMaxKey() string {
	if this.maxh.Len() == 0 {
		return ""
	}
	return this.maxh[0].k
}

func (this *AllOne) GetMinKey() string {
	if this.minh.Len() == 0 {
		return ""
	}
	return this.minh[0].k
}

type Pair struct {
	k string
	v int
}

type MaxHeap []*Pair

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].v > h[j].v }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(*Pair)) }
func (h *MaxHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}

type MinHeap []*Pair

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].v < h[j].v }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*Pair)) }
func (h *MinHeap) Pop() interface{} {
	i := h.Len() - 1
	out := (*h)[i]
	*h = (*h)[:i]
	return out
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
