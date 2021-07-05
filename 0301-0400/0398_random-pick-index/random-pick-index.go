package main

import "math/rand"

// Reservoir Sampling

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (this *Solution) Pick(target int) int {
	out := -1
	idx := 0
	for i, num := range this.nums {
		if num != target {
			continue
		}
		idx++
		if rand.Intn(idx)+1 == idx {
			out = i
		}
	}
	return out
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
