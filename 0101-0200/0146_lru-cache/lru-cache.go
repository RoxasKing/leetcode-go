package main

/*
  运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
  获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
  写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

  进阶: 你是否可以在 O(1) 时间复杂度内完成这两种操作？
*/

type LRUCache struct {
	head     *cacheNode
	tail     *cacheNode
	dict     map[int]*cacheNode
	capacity int
}

func Constructor(capacity int) LRUCache {
	head, tail := newCacheNode(-1<<31, -1<<31), newCacheNode(-1<<31, -1<<31)
	head.next, tail.prev = tail, head
	return LRUCache{
		head:     head,
		tail:     tail,
		dict:     make(map[int]*cacheNode, capacity),
		capacity: capacity,
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
	if len(this.dict) == this.capacity {
		first := this.head.next
		this.removeNode(first)
		delete(this.dict, first.key)
	}
	node := newCacheNode(key, value)
	this.dict[key] = node
	this.insertToTail(node)
}

type cacheNode struct {
	key  int
	val  int
	prev *cacheNode
	next *cacheNode
}

func newCacheNode(key, value int) *cacheNode {
	return &cacheNode{key: key, val: value}
}

func (l *LRUCache) insertToTail(node *cacheNode) {
	tail := l.tail
	node.prev = tail.prev
	tail.prev.next = node
	node.next = tail
	tail.prev = node
}

func (l *LRUCache) removeNode(node *cacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
