package main

import "math/rand"

type RandomizedSet struct {
	hash map[int]int
	list []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		hash: make(map[int]int),
		list: []int{},
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hash[val]; ok {
		return false
	}
	this.hash[val] = len(this.list)
	this.list = append(this.list, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	var index int
	var ok bool
	if index, ok = this.hash[val]; !ok {
		return false
	}
	last := len(this.list) - 1
	lastVal := this.list[last]
	this.hash[lastVal] = index
	this.list[index], this.list[last] = this.list[last], this.list[index]
	this.list = this.list[:last]
	delete(this.hash, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if len(this.list) == 0 {
		return -1
	}
	randIndex := rand.Intn(len(this.list))
	return this.list[randIndex]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
