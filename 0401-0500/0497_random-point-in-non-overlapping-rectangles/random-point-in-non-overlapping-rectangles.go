package main

import (
	"math/rand"
	"sort"
)

// Difficulty:
// Medium

// Tags:
// Randomized
// Binary Search

type area struct{ x, y, i, n int }

type Solution struct {
	n int
	a []area
}

func Constructor(rects [][]int) Solution {
	idx := 0
	arr := []area{}
	for _, e := range rects {
		a, b, c, d := e[0], e[1], e[2], e[3]
		arr = append(arr, area{a, b, idx, d - b + 1})
		idx += (c - a + 1) * (d - b + 1)
	}
	return Solution{idx, arr}
}

func (this *Solution) Pick() []int {
	pos := rand.Intn(this.n)
	i := sort.Search(len(this.a), func(i int) bool { return this.a[i].i > pos }) - 1
	k, n := pos-this.a[i].i, this.a[i].n
	return []int{this.a[i].x + k/n, this.a[i].y + k%n}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
