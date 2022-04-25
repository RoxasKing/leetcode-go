package main

// Difficulty:
// Easy

// Tags:
// Design

type MyHashMap struct {
	b [1000]*bucket
}

type bucket struct{ h [1001]int }

func Constructor() MyHashMap {
	return MyHashMap{}
}

func (this *MyHashMap) Put(key int, value int) {
	i, j := key/1001, key%1001
	if this.b[i] == nil {
		this.b[i] = &bucket{}
		for j := 0; j < 1001; j++ {
			this.b[i].h[j] = -1
		}
	}
	this.b[i].h[j] = value
}

func (this *MyHashMap) Get(key int) int {
	i, j := key/1001, key%1001
	if this.b[i] == nil {
		return -1
	}
	return this.b[i].h[j]
}

func (this *MyHashMap) Remove(key int) {
	i, j := key/1001, key%1001
	if this.b[i] == nil {
		return
	}
	this.b[i].h[j] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
