package main

import "container/heap"

// Difficulty:
// Hard

// Tags:
// Hash
// Priority Queue

type AllOne struct {
	freq       map[string]int
	maxh, minh *ps
}

func Constructor() AllOne {
	return AllOne{freq: map[string]int{}, maxh: &ps{f: bigger}, minh: &ps{f: smaller}}
}

func (this *AllOne) Inc(key string) {
	this.freq[key]++
	for ; this.maxh.Len() > 0 && this.maxh.Top().v != this.freq[this.maxh.Top().k]; heap.Pop(this.maxh) {
	}
	for ; this.minh.Len() > 0 && this.minh.Top().v != this.freq[this.minh.Top().k]; heap.Pop(this.minh) {
	}
	heap.Push(this.maxh, &pair{key, this.freq[key]})
	heap.Push(this.minh, &pair{key, this.freq[key]})
}

func (this *AllOne) Dec(key string) {
	this.freq[key]--
	for ; this.maxh.Len() > 0 && this.maxh.Top().v != this.freq[this.maxh.Top().k]; heap.Pop(this.maxh) {
	}
	for ; this.minh.Len() > 0 && this.minh.Top().v != this.freq[this.minh.Top().k]; heap.Pop(this.minh) {
	}
	if this.freq[key] == 0 {
		return
	}
	heap.Push(this.maxh, &pair{key, this.freq[key]})
	heap.Push(this.minh, &pair{key, this.freq[key]})
}

func (this *AllOne) GetMaxKey() string {
	if this.maxh.Len() == 0 {
		return ""
	}
	return this.maxh.Top().k
}

func (this *AllOne) GetMinKey() string {
	if this.minh.Len() == 0 {
		return ""
	}
	return this.minh.Top().k
}

type pair struct {
	k string
	v int
}

type ps struct {
	a []*pair
	f func(i, j int) bool
}

func (s ps) Len() int            { return len(s.a) }
func (s ps) Less(i, j int) bool  { return s.f(s.a[i].v, s.a[j].v) }
func (s ps) Swap(i, j int)       { s.a[i], s.a[j] = s.a[j], s.a[i] }
func (s *ps) Push(x interface{}) { s.a = append(s.a, x.(*pair)) }
func (s *ps) Pop() interface{}   { i := s.Len() - 1; o := s.a[i]; s.a = s.a[:i]; return o }
func (s ps) Top() *pair          { return s.a[0] }

func bigger(a, b int) bool  { return a > b }
func smaller(a, b int) bool { return a < b }

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
