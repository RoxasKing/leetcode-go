package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Hash
// Binary Search

type TimeMap struct {
	h map[string][]*pair
}

type pair struct {
	t int
	v string
}

func Constructor() TimeMap {
	return TimeMap{map[string][]*pair{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.h[key] = append(this.h[key], &pair{v: value, t: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	a := this.h[key]
	if i := sort.Search(len(a), func(i int) bool { return a[i].t > timestamp }) - 1; i >= 0 {
		return a[i].v
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
