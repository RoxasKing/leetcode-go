package main

type MyHashMap struct {
	b [1001]*bucket
}

type bucket struct {
	h     [1001]int
	valid [1001]bool
	cnt   int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		b: [1001]*bucket{},
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	i, j := key/1000, key%1000
	if this.b[i] == nil {
		this.b[i] = &bucket{}
	}
	this.b[i].h[j] = value
	this.b[i].valid[j] = true
	this.b[i].cnt++
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	i, j := key/1000, key%1000
	if this.b[i] == nil || !this.b[i].valid[j] {
		return -1
	}
	return this.b[i].h[j]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	i, j := key/1000, key%1000
	if this.b[i] == nil || !this.b[i].valid[j] {
		return
	}
	this.b[i].valid[j] = false
	this.b[i].cnt--
	if this.b[i].cnt == 0 {
		this.b[i] = nil
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
