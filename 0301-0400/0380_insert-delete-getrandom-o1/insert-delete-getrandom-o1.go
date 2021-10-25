package main

import "math/rand"

// Difficulty:
// Medium

// Tags:
// Hash

type RandomizedSet struct {
	idx map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		idx: map[int]int{},
		arr: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.idx[val]; ok {
		return false
	}
	this.idx[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.idx[val]; !ok {
		return false
	}
	idx := this.idx[val]
	end := len(this.arr) - 1
	this.idx[this.arr[end]] = idx
	this.arr[idx], this.arr[end] = this.arr[end], this.arr[idx]
	delete(this.idx, val)
	this.arr = this.arr[:end]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
