package main

/*
  Design a HashSet without using any built-in hash table libraries.

  Implement MyHashSet class:
    1. void add(key) Inserts the value key into the HashSet.
    2. bool contains(key) Returns whether the value key exists in the HashSet or not.
    3. void remove(key) Removes the value key in the HashSet. If key does not exist in the HashSet, do nothing.

  Example 1:
    Input
      ["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
      [[], [1], [2], [1], [3], [2], [2], [2], [2]]
    Output
      [null, null, null, true, false, null, true, null, false]
    Explanation
      MyHashSet myHashSet = new MyHashSet();
      myHashSet.add(1);      // set = [1]
      myHashSet.add(2);      // set = [1, 2]
      myHashSet.contains(1); // return True
      myHashSet.contains(3); // return False, (not found)
      myHashSet.add(2);      // set = [1, 2]
      myHashSet.contains(2); // return True
      myHashSet.remove(2);   // set = [1]
      myHashSet.contains(2); // return False, (already removed)

  Constraints:
    1. 0 <= key <= 10^6
    2. At most 10^4 calls will be made to add, remove, and contains.

  Follow up: Could you solve the problem without using the built-in HashSet library?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/design-hashset
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
