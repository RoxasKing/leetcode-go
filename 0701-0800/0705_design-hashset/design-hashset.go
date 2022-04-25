package main

// Difficulty:
// Easy

// Tags:
// Design
// Bit Manipulation

type MyHashSet struct {
	bitset []int
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	idx, bit := key/64, key%64
	for i := len(this.bitset) - 1; i < idx; i++ {
		this.bitset = append(this.bitset, 0)
	}
	this.bitset[idx] |= 1 << bit
}

func (this *MyHashSet) Remove(key int) {
	idx, bit := key/64, key%64
	if len(this.bitset) <= idx {
		return
	}
	this.bitset[idx] &= ^(1 << bit)
}

func (this *MyHashSet) Contains(key int) bool {
	idx, bit := key/64, key%64
	if len(this.bitset) <= idx {
		return false
	}
	return this.bitset[idx]&(1<<bit) != 0
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
