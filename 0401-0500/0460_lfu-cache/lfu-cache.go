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

type LFUCache struct {
	cacheMap  map[int]*cacheNode
	frequency map[int]*doublyLList
	minFreq   int
	capacity  int
	size      int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cacheMap:  make(map[int]*cacheNode),
		frequency: make(map[int]*doublyLList),
		capacity:  capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	node, ok := this.cacheMap[key]
	if !ok {
		return -1
	}
	this.increaseFreq(node)
	return node.val
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if node, ok := this.cacheMap[key]; ok {
		this.cacheMap[key].val = value
		this.increaseFreq(node)
		return
	}
	if this.size == this.capacity {
		this.removeMin()
	}
	this.addNode(&cacheNode{
		key:  key,
		val:  value,
		freq: 1,
	})
}

func (this *LFUCache) increaseFreq(node *cacheNode) {
	this.frequency[node.freq].remove(node)
	if node.freq == this.minFreq && this.frequency[node.freq].size == 0 {
		this.minFreq++
	}
	node.freq++
	if this.frequency[node.freq] == nil {
		this.frequency[node.freq] = newDoubleList()
	}
	this.frequency[node.freq].pushBack(node)
}

func (this *LFUCache) addNode(node *cacheNode) {
	if this.frequency[1] == nil {
		this.frequency[1] = newDoubleList()
	}
	this.frequency[1].pushBack(node)
	this.cacheMap[node.key] = node
	this.minFreq = 1
	this.size++
}

func (c *LFUCache) removeMin() {
	toRemove := c.frequency[c.minFreq].front()
	c.frequency[c.minFreq].remove(toRemove)
	delete(c.cacheMap, toRemove.key)
	if c.frequency[c.minFreq].size == 0 {
		c.minFreq++
	}
	c.size--
}

type cacheNode struct {
	key  int
	val  int
	freq int
	prev *cacheNode
	next *cacheNode
}

type doublyLList struct {
	head *cacheNode
	tail *cacheNode
	size int
}

func newDoubleList() *doublyLList {
	head := &cacheNode{}
	tail := &cacheNode{prev: head}
	head.next = tail
	return &doublyLList{
		head: head,
		tail: tail,
	}
}

func (l *doublyLList) pushBack(node *cacheNode) {
	node.prev, node.next = l.tail.prev, l.tail
	node.prev.next, node.next.prev = node, node
	l.size++
}

func (l *doublyLList) front() *cacheNode {
	return l.head.next
}

func (l *doublyLList) remove(node *cacheNode) {
	node.prev.next, node.next.prev = node.next, node.prev
	l.size--
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
