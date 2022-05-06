package main

import "sort"

// Difficulty:
// Easy

// Tags:
// Binary Search

type RecentCounter struct {
	a []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.a = append(this.a, t)
	i := sort.Search(len(this.a), func(i int) bool { return this.a[i] >= t-3000 })
	return len(this.a) - i
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
