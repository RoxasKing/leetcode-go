package main

// Tags:
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
