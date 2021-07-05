package main

import (
	"math/rand"
	"sort"
)

// Tags:
// Binary Search

type Solution struct {
	n   int
	arr [][]int
}

func Constructor(rects [][]int) Solution {
	arr := [][]int{}
	start := 0
	for _, rect := range rects {
		a, b, c, d := rect[0], rect[1], rect[2], rect[3]
		arr = append(arr, []int{start, d - b + 1, a, b})
		start += (c - a + 1) * (d - b + 1)
	}
	return Solution{
		n:   start,
		arr: arr,
	}
}

func (this *Solution) Pick() []int {
	idx := rand.Intn(this.n)
	i := sort.Search(len(this.arr), func(i int) bool { return this.arr[i][0] > idx }) - 1
	extra := idx - this.arr[i][0]
	n, x, y := this.arr[i][1], this.arr[i][2], this.arr[i][3]
	return []int{x + extra/n, y + extra%n}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
