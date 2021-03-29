package main

/*
  Design a HashMap without using any built-in hash table libraries.

  Implement the MyHashMap class:
    1. MyHashMap() initializes the object with an empty map.
    2. void put(int key, int value) inserts a (key, value) pair into the HashMap. If the key already exists in the map, update the corresponding value.
    3. int get(int key) returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
    4. void remove(key) removes the key and its corresponding value if the map contains the mapping for the key.

  Example 1:
    Input
      ["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
      [[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
    Output
      [null, null, null, 1, -1, null, 1, null, -1]
    Explanation
      MyHashMap myHashMap = new MyHashMap();
      myHashMap.put(1, 1); // The map is now [[1,1]]
      myHashMap.put(2, 2); // The map is now [[1,1], [2,2]]
      myHashMap.get(1);    // return 1, The map is now [[1,1], [2,2]]
      myHashMap.get(3);    // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
      myHashMap.put(2, 1); // The map is now [[1,1], [2,1]] (i.e., update the existing value)
      myHashMap.get(2);    // return 1, The map is now [[1,1], [2,1]]
      myHashMap.remove(2); // remove the mapping for 2, The map is now [[1,1]]
      myHashMap.get(2);    // return -1 (i.e., not found), The map is now [[1,1]]

  Constraints:
    1. 0 <= key, value <= 10^6
    2. At most 10^4 calls will be made to put, get, and remove.

  Follow up: Please do not use the built-in HashMap library.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/design-hashmap
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
