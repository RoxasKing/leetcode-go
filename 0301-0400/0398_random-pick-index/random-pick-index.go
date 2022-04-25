package main

import "math/rand"

// Difficulty:
// Medium

// Tags:
// Reservoir Sampling

type Solution struct {
	a []int
	i int
}

func Constructor(nums []int) Solution {
	return Solution{nums, -1}
}

func (this *Solution) Pick(target int) int {
	o, i := -1, 0
	for o == -1 {
		this.i = (this.i + 1) % len(this.a)
		if this.a[this.i] != target {
			continue
		}
		if i++; rand.Intn(i)+1 == i {
			o = this.i
		}
	}
	return o
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
