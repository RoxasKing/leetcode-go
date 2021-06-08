package main

import "container/heap"

/*
  Design a data structure to store the strings' count with the ability to return the strings with minimum and maximum counts.

  Implement the AllOne class:

    1. AllOne() Initializes the object of the data structure.
    2. inc(String key) Increments the count of the string key by 1. If key does not exist in the data structure, insert it with count 1.
    3. dec(String key) Decrements the count of the string key by 1. If the count of key is 0 after the decrement, remove it from the data structure. It is guaranteed that key exists in the data structure before the decrement.
    4. getMaxKey() Returns one of the keys with the maximal count. If no element exists, return an empty string "".
    5. getMinKey() Returns one of the keys with the minimum count. If no element exists, return an empty string "".

  Example 1:
    Input
      ["AllOne", "inc", "inc", "getMaxKey", "getMinKey", "inc", "getMaxKey", "getMinKey"]
      [[], ["hello"], ["hello"], [], [], ["leet"], [], []]
    Output
      [null, null, null, "hello", "hello", null, "hello", "leet"]

  Explanation
    AllOne allOne = new AllOne();
    allOne.inc("hello");
    allOne.inc("hello");
    allOne.getMaxKey(); // return "hello"
    allOne.getMinKey(); // return "hello"
    allOne.inc("leet");
    allOne.getMaxKey(); // return "hello"
    allOne.getMinKey(); // return "leet"

  Constraints:
    1. 1 <= key.length <= 10
    2. key consists of lowercase English letters.
    3. It is guaranteed that for each call to dec, key is existing in the data structure.
    4. At most 5 * 104 calls will be made to inc, dec, getMaxKey, and getMinKey.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/all-oone-data-structure
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue

type AllOne struct {
	h    map[string]int
	maxh MaxHeap
	minh MinHeap
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		h:    map[string]int{},
		maxh: MaxHeap{},
		minh: MinHeap{},
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	this.h[key]++
	for this.maxh.Len() > 0 && this.maxh[0].cnt != this.h[this.maxh[0].key] {
		heap.Pop(&this.maxh)
	}
	for this.minh.Len() > 0 && this.minh[0].cnt != this.h[this.minh[0].key] {
		heap.Pop(&this.minh)
	}
	heap.Push(&this.maxh, &Pair{key: key, cnt: this.h[key]})
	heap.Push(&this.minh, &Pair{key: key, cnt: this.h[key]})
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	this.h[key]--
	for this.maxh.Len() > 0 && this.maxh[0].cnt != this.h[this.maxh[0].key] {
		heap.Pop(&this.maxh)
	}
	for this.minh.Len() > 0 && this.minh[0].cnt != this.h[this.minh[0].key] {
		heap.Pop(&this.minh)
	}
	if this.h[key] == 0 {
		delete(this.h, key)
		return
	}
	heap.Push(&this.maxh, &Pair{key: key, cnt: this.h[key]})
	heap.Push(&this.minh, &Pair{key: key, cnt: this.h[key]})
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.maxh.Len() > 0 {
		return this.maxh[0].key
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.minh.Len() > 0 {
		return this.minh[0].key
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */

type Pair struct {
	key string
	cnt int
}

type MaxHeap []*Pair

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].cnt > h[j].cnt }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(*Pair)) }
func (h *MaxHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}

type MinHeap []*Pair

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].cnt < h[j].cnt }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*Pair)) }
func (h *MinHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
