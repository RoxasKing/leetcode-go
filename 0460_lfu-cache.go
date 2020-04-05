package leetcode

/*
  设计并实现最不经常使用（LFU）缓存的数据结构。它应该支持以下操作：get 和 put。
    get(key) - 如果键存在于缓存中，则获取键的值（总是正数），否则返回 -1。
    put(key, value) - 如果键不存在，请设置或插入值。当缓存达到其容量时，它应该在插入新项目之前，使最不经常使用的项目无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，最近最少使用的键将被去除。
  进阶：
  你是否可以在 O(1) 时间复杂度内执行两项操作？
*/

// LFUCache ...
type LFUCache struct {
	cacheMap  map[int]*cacheNode
	frequency map[int]*doubleList
	minFreq   int
	capacity  int
	size      int
}

// NewLFUCache ...
func NewLFUCache(capacity int) LFUCache {
	return LFUCache{
		cacheMap:  make(map[int]*cacheNode),
		frequency: make(map[int]*doubleList),
		capacity:  capacity,
	}
}

// Get ...
func (c *LFUCache) Get(key int) int {
	node, ok := c.cacheMap[key]
	if !ok {
		return -1
	}
	c.incrementFreq(node)
	return node.val
}

// Put ...
func (c *LFUCache) Put(key int, value int) {
	if c.capacity == 0 {
		return
	}
	if node, ok := c.cacheMap[key]; ok {
		c.cacheMap[key].val = value
		c.incrementFreq(node)
		return
	}
	if c.size == c.capacity {
		c.removeMin()
	}
	c.addNode(&cacheNode{
		key:  key,
		val:  value,
		freq: 1,
	})
}

func (c *LFUCache) incrementFreq(node *cacheNode) {
	c.frequency[node.freq].remove(node)
	if node.freq == c.minFreq && c.frequency[node.freq].size == 0 {
		c.minFreq++
	}
	node.freq++
	if c.frequency[node.freq] == nil {
		c.frequency[node.freq] = newDoubleList()
	}
	c.frequency[node.freq].pushBack(node)
}

func (c *LFUCache) addNode(node *cacheNode) {
	if c.frequency[1] == nil {
		c.frequency[1] = newDoubleList()
	}
	c.frequency[1].pushBack(node)
	c.cacheMap[node.key] = node
	c.minFreq = 1
	c.size++
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

type doubleList struct {
	head *cacheNode
	tail *cacheNode
	size int
}

func newDoubleList() *doubleList {
	head := &cacheNode{}
	tail := &cacheNode{prev: head}
	head.next = tail
	return &doubleList{
		head: head,
		tail: tail,
	}
}

func (l *doubleList) pushBack(node *cacheNode) {
	node.prev, node.next = l.tail.prev, l.tail
	node.prev.next, node.next.prev = node, node
	l.size++
}

func (l *doubleList) front() *cacheNode {
	return l.head.next
}

func (l *doubleList) remove(node *cacheNode) {
	node.prev.next, node.next.prev = node.next, node.prev
	l.size--
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
