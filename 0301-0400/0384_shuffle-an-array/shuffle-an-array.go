package main

import "math/rand"

type Solution struct {
	origin []int
	array  []int
}

func Constructor(nums []int) Solution {
	array := make([]int, len(nums))
	copy(array, nums)
	return Solution{
		origin: nums,
		array:  array,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	this.array = this.origin
	this.origin = make([]int, len(this.array))
	copy(this.origin, this.array)
	return this.origin
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i := range this.array {
		tmp := rand.Intn(len(this.array))
		this.array[i], this.array[tmp] = this.array[tmp], this.array[i]
	}
	return this.array
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
