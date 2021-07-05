package main

type MyHashSet struct {
	bitset []uint64
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		bitset: []uint64{},
	}
}

func (this *MyHashSet) Add(key int) {
	bit := key % 64
	idx := key / 64
	for i := len(this.bitset); i <= idx; i++ {
		this.bitset = append(this.bitset, 0)
	}
	this.bitset[idx] |= 1 << bit
}

func (this *MyHashSet) Remove(key int) {
	bit := key % 64
	idx := key / 64
	if idx >= len(this.bitset) {
		return
	}
	this.bitset[idx] &= ^(1 << bit)
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	bit := key % 64
	idx := key / 64
	if idx >= len(this.bitset) {
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
