package main

import "container/heap"

// Difficulty:
// Hard

// Tags:
// Hash
// Priority Queue

type AllOne struct {
	freq       map[string]int
	maxh, minh ph
}

func Constructor() AllOne {
	return AllOne{
		freq: map[string]int{},
		maxh: ph{f: func(a, b int) bool { return a > b }},
		minh: ph{f: func(a, b int) bool { return a < b }},
	}
}

func (this *AllOne) Inc(key string) {
	this.freq[key]++
	for ; this.maxh.Len() > 0 && this.maxh.a[0].v != this.freq[this.maxh.a[0].k]; heap.Pop(&this.maxh) {
	}
	for ; this.minh.Len() > 0 && this.minh.a[0].v != this.freq[this.minh.a[0].k]; heap.Pop(&this.minh) {
	}
	heap.Push(&this.maxh, &pair{key, this.freq[key]})
	heap.Push(&this.minh, &pair{key, this.freq[key]})
}

func (this *AllOne) Dec(key string) {
	this.freq[key]--
	for ; this.maxh.Len() > 0 && this.maxh.a[0].v != this.freq[this.maxh.a[0].k]; heap.Pop(&this.maxh) {
	}
	for ; this.minh.Len() > 0 && this.minh.a[0].v != this.freq[this.minh.a[0].k]; heap.Pop(&this.minh) {
	}
	if this.freq[key] == 0 {
		return
	}
	heap.Push(&this.maxh, &pair{key, this.freq[key]})
	heap.Push(&this.minh, &pair{key, this.freq[key]})
}

func (this *AllOne) GetMaxKey() string {
	if this.maxh.Len() == 0 {
		return ""
	}
	return this.maxh.a[0].k
}

func (this *AllOne) GetMinKey() string {
	if this.minh.Len() == 0 {
		return ""
	}
	return this.minh.a[0].k
}

type pair struct {
	k string
	v int
}

type ph struct {
	a []*pair
	f func(i, j int) bool
}

func (h ph) Len() int            { return len(h.a) }
func (h ph) Less(i, j int) bool  { return h.f(h.a[i].v, h.a[j].v) }
func (h ph) Swap(i, j int)       { h.a[i], h.a[j] = h.a[j], h.a[i] }
func (h *ph) Push(x interface{}) { h.a = append(h.a, x.(*pair)) }
func (h *ph) Pop() interface{}   { i := h.Len() - 1; o := (h.a)[i]; h.a = (h.a)[:i]; return o }

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
