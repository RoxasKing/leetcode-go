package main

import (
	"math/rand"
	"sort"
)

// Tags:
// Binary Search

type Solution struct {
	n int
	b []int
}

func Constructor(n int, blacklist []int) Solution {
	sort.Ints(blacklist)
	return Solution{
		n: n,
		b: blacklist,
	}
}

func (this *Solution) Pick() int {
	x := rand.Intn(this.n - len(this.b))
	l, r := 0, len(this.b)-1
	for l < r {
		m := (l + r + 1) >> 1
		if this.b[m]-m <= x {
			l = m
		} else {
			r = m - 1
		}
	}
	if l == r && this.b[l]-l <= x {
		return x + l + 1
	}
	return x
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
