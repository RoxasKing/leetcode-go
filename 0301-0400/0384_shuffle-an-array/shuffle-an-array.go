package main

import "math/rand"

// Difficulty:
// Medium

// Tags:
// Math
// Radomized

type Solution struct {
	arr, cur []int
}

func Constructor(nums []int) Solution {
	cur := make([]int, len(nums))
	copy(cur, nums)
	return Solution{arr: nums, cur: cur}
}

func (this *Solution) Reset() []int {
	copy(this.cur, this.arr)
	return this.arr
}

func (this *Solution) Shuffle() []int {
	for i := range this.cur {
		j := rand.Intn(len(this.cur))
		this.cur[i], this.cur[j] = this.cur[j], this.cur[i]
	}
	return this.cur
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
