package main

import "math/rand"

// Difficulty:
// Medium

// Tags:
// Hash

type RandomizedSet struct {
	h map[int]int
	a []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{h: map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.h[val]; ok {
		return false
	}
	i := len(this.a)
	this.a = append(this.a, val)
	this.h[val] = i
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if i, ok := this.h[val]; !ok {
		return false
	} else if i != len(this.a)-1 {
		x := this.a[len(this.a)-1]
		j := this.h[x]
		this.a[i], this.a[j] = this.a[j], this.a[i]
		this.h[x] = i
	}
	this.a = this.a[:len(this.a)-1]
	delete(this.h, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.a[rand.Intn(len(this.a))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
