package main

import "math/rand"

// Difficulty:
// Medium

// Tags:
// Randomized

type Solution struct {
	m, n, total int
	h           map[int]int
}

func Constructor(m int, n int) Solution {
	return Solution{m: m, n: n, total: m * n, h: map[int]int{}}
}

func (this *Solution) Flip() []int {
	i := rand.Intn(this.total)
	this.total--
	var out []int
	if j, ok := this.h[i]; ok {
		out = []int{j / this.n, j % this.n}
	} else {
		out = []int{i / this.n, i % this.n}
	}
	if j, ok := this.h[this.total]; ok {
		this.h[i] = j
	} else {
		this.h[i] = this.total
	}
	return out
}

func (this *Solution) Reset() {
	this.total = this.m * this.n
	this.h = map[int]int{}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
