package main

import (
	"math/rand"
	"sort"
)

// Tags:
// Prefix Sum + Binary Search

type Solution struct {
	pSum []int
}

func Constructor(w []int) Solution {
	n := len(w)
	pSum := make([]int, n+1)
	for i := range w {
		pSum[i+1] = pSum[i] + w[i]
	}
	return Solution{pSum: pSum}
}

func (this *Solution) PickIndex() int {
	num := rand.Intn(this.pSum[len(this.pSum)-1]) + 1
	i := sort.Search(len(this.pSum), func(i int) bool { return this.pSum[i] >= num })
	return i - 1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
