package main

import (
	"math/rand"
	"sort"
)

// Tags:
// Prefix Sum
// Binary Search

type Solution struct {
	p []int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{p: w}
}

func (this *Solution) PickIndex() int {
	x := rand.Intn(this.p[len(this.p)-1]) + 1
	return sort.SearchInts(this.p, x)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
