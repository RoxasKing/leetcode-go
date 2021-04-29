package main

/*
  Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

  Implement the LRUCache class:
    1. LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
    2. int get(int key) Return the value of the key if the key exists, otherwise return -1.
    3. void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.

  Follow up:
    Could you do get and put in O(1) time complexity?

  Example 1:
  Input
    ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
    [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
  Output
    [null, null, null, 1, null, -1, null, -1, 3, 4]
  Explanation
    LRUCache lRUCache = new LRUCache(2);
    lRUCache.put(1, 1); // cache is {1=1}
    lRUCache.put(2, 2); // cache is {1=1, 2=2}
    lRUCache.get(1);    // return 1
    lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
    lRUCache.get(2);    // returns -1 (not found)
    lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
    lRUCache.get(1);    // return -1 (not found)
    lRUCache.get(3);    // return 3
    lRUCache.get(4);    // return 4

  Constraints:
    1. 1 <= capacity <= 3000
    2. 0 <= key <= 3000
    3. 0 <= value <= 10^4
    4. At most 3 * 10^4 calls will be made to get and put.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/lru-cache
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Design
type LRUCache struct {
	head *cacheNode
	tail *cacheNode
	dict map[int]*cacheNode
	cap  int
}

func Constructor(capacity int) LRUCache {
	head, tail := newCacheNode(-1, -1), newCacheNode(-1, -1)
	head.next = tail
	tail.prev = head
	return LRUCache{
		head: head,
		tail: tail,
		dict: make(map[int]*cacheNode),
		cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.dict[key]; ok {
		this.removeNode(node)
		this.insertToTail(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.dict[key]; ok {
		this.removeNode(node)
		this.insertToTail(node)
		node.val = value
		return
	}
	if len(this.dict) == this.cap {
		first := this.head.next
		this.removeNode(first)
		delete(this.dict, first.key)
	}
	node := newCacheNode(key, value)
	this.dict[key] = node
	this.insertToTail(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func (this *LRUCache) insertToTail(node *cacheNode) {
	tail := this.tail
	node.prev = tail.prev
	tail.prev.next = node
	node.next = tail
	tail.prev = node
}

func (this *LRUCache) removeNode(node *cacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

type cacheNode struct {
	key  int
	val  int
	prev *cacheNode
	next *cacheNode
}

func newCacheNode(key int, val int) *cacheNode {
	return &cacheNode{key: key, val: val}
}
