package main

// Tags:
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
