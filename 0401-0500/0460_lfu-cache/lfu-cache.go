package main

/*
  Design and implement a data structure for a Least Frequently Used (LFU) cache.

  Implement the LFUCache class:
    1. LFUCache(int capacity) Initializes the object with the capacity of the data structure.
    2. int get(int key) Gets the value of the key if the key exists in the cache. Otherwise, returns -1.
    3. void put(int key, int value) Update the value of the key if present, or inserts the key if not already present. When the cache reaches its capacity, it should invalidate and remove the least frequently used key before inserting a new item. For this problem, when there is a tie (i.e., two or more keys with the same frequency), the least recently used key would be invalidated.
  To determine the least frequently used key, a use counter is maintained for each key in the cache. The key with the smallest use counter is the least frequently used key.

  When a key is first inserted into the cache, its use counter is set to 1 (due to the put operation). The use counter for a key in the cache is incremented either a get or put operation is called on it.

  Example 1:
    Input
      ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
      [[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
    Output
      [null, null, null, 1, null, -1, 3, null, -1, 3, 4]
    Explanation
      // cnt(x) = the use counter for key x
      // cache=[] will show the last used order for tiebreakers (leftmost element is  most recent)
      LFUCache lfu = new LFUCache(2);
      lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
      lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
      lfu.get(1);      // return 1
                       // cache=[1,2], cnt(2)=1, cnt(1)=2
      lfu.put(3, 3);   // 2 is the LFU key because cnt(2)=1 is the smallest, invalidate 2.
                       // cache=[3,1], cnt(3)=1, cnt(1)=2
      lfu.get(2);      // return -1 (not found)
      lfu.get(3);      // return 3
                       // cache=[3,1], cnt(3)=2, cnt(1)=2
      lfu.put(4, 4);   // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1.
                       // cache=[4,3], cnt(4)=1, cnt(3)=2
      lfu.get(1);      // return -1 (not found)
      lfu.get(3);      // return 3
                       // cache=[3,4], cnt(4)=1, cnt(3)=3
      lfu.get(4);      // return 4
                       // cache=[3,4], cnt(4)=2, cnt(3)=3

  Constraints:
    1. 0 <= capacity, key, value <= 10^4
    2. At most 10^5 calls will be made to get and put.

  Follow up: Could you do both operations in O(1) time complexity?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/lfu-cache
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Doubly Linked List + Hash

type LFUCache struct {
	nodeAddr map[int]*cacheNode
	freqList map[int]*doublyLL
	minFreq  int
	cap, len int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		nodeAddr: make(map[int]*cacheNode),
		freqList: make(map[int]*doublyLL),
		cap:      capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.nodeAddr[key]; ok {
		this.incrFreq(node)
		return node.val
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}

	if node, ok := this.nodeAddr[key]; ok {
		node.val = value
		this.incrFreq(node)
		return
	}

	if this.len == this.cap {
		del := this.freqList[this.minFreq].peekHead()
		this.freqList[this.minFreq].remove(del)
		delete(this.nodeAddr, del.key)
		this.len--
	}

	node := newCacheNode(key, value)
	node.freq = 1
	this.nodeAddr[key] = node
	if this.freqList[1] == nil {
		this.freqList[1] = newDoublyLL()
	}
	this.freqList[1].pushToTail(node)
	this.minFreq = 1
	this.len++
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func (this *LFUCache) incrFreq(node *cacheNode) {
	this.freqList[node.freq].remove(node)
	if this.minFreq == node.freq && this.freqList[node.freq].len == 0 {
		this.minFreq++
	}
	node.freq++
	if this.freqList[node.freq] == nil {
		this.freqList[node.freq] = newDoublyLL()
	}
	this.freqList[node.freq].pushToTail(node)
}

type doublyLL struct {
	head, tail *cacheNode
	len        int
}

func newDoublyLL() *doublyLL {
	head := newCacheNode(-1, -1)
	tail := newCacheNode(-1, -1)
	head.next = tail
	tail.prev = head
	return &doublyLL{head: head, tail: tail}
}

func (ll *doublyLL) pushToTail(node *cacheNode) {
	node.prev = ll.tail.prev
	node.next = ll.tail
	node.prev.next = node
	node.next.prev = node
	ll.len++
}

func (ll *doublyLL) remove(node *cacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	ll.len--
}

func (ll doublyLL) peekHead() *cacheNode {
	return ll.head.next
}

type cacheNode struct {
	key, val, freq int
	prev, next     *cacheNode
}

func newCacheNode(key, val int) *cacheNode {
	return &cacheNode{key: key, val: val}
}
