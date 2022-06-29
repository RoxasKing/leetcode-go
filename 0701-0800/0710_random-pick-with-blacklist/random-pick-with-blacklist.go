package main

import (
	"math/rand"
	"sort"
)

// Difficulty:
// Hard

// Tags:
// Binary Search

type Solution struct {
	n int
	b []int
}

func Constructor(n int, blacklist []int) Solution {
	sort.Ints(blacklist)
	return Solution{n, blacklist}
}

func (this *Solution) Pick() int {
	o := rand.Intn(this.n - len(this.b))
	if len(this.b) == 0 {
		return o
	}
	l, r := 0, len(this.b)-1
	for l < r {
		m := (l + r + 1) >> 1
		if this.b[m]-m <= o {
			l = m
		} else {
			r = m - 1
		}
	}
	if this.b[l]-l <= o {
		o += l + 1
	}
	return o
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
