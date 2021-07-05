package main

import (
	"math/rand"
	"sort"
)

// Tags:
// Hash
type RandomizedCollection struct {
	idxs map[int][]int
	list []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		idxs: make(map[int][]int),
		list: []int{},
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	success := false
	if _, ok := this.idxs[val]; !ok {
		success = true
	}
	this.idxs[val] = append(this.idxs[val], len(this.list))
	this.list = append(this.list, val)
	return success
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	if arr, ok := this.idxs[val]; !ok || len(arr) == 0 {
		return false
	}
	idx := this.idxs[val][len(this.idxs[val])-1]
	this.idxs[val] = this.idxs[val][:len(this.idxs[val])-1]
	tail := len(this.list) - 1
	tailVal := this.list[tail]
	if tailVal == val {
		this.list = this.list[:tail]
		return true
	}
	this.list[tail], this.list[idx] = this.list[idx], this.list[tail]
	this.list = this.list[:tail]
	this.idxs[tailVal] = append(this.idxs[tailVal][:len(this.idxs[tailVal])-1], idx)
	sort.Ints(this.idxs[tailVal])
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
