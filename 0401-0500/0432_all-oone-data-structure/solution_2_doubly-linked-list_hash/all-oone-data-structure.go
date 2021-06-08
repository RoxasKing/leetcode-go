package main

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

// Important!

// Doubly Linked List
// Hash

type AllOne struct {
	nodeAddr   map[string]*dataNode
	freqList   map[int]*doublyLL
	head, tail *doublyLL
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	head := newDoublyLL()
	tail := newDoublyLL()
	head.next = tail
	tail.prev = head
	return AllOne{
		nodeAddr: map[string]*dataNode{},
		freqList: map[int]*doublyLL{},
		head:     head,
		tail:     tail,
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	node, ok := this.nodeAddr[key]

	if !ok {
		node = newDataNode(key)
		this.nodeAddr[key] = node
		node.freq = 1
		if this.freqList[1] == nil {
			this.freqList[1] = newDoublyLL()
			this.head.next.prev = this.freqList[1]
			this.freqList[1].next = this.head.next
			this.head.next = this.freqList[1]
			this.freqList[1].prev = this.head
		}
		this.freqList[1].pushToTail(node)
		return
	}

	ll := this.freqList[node.freq]
	node.freq++

	// just change ll's freq
	if ll.len == 1 && this.freqList[node.freq] == nil {
		delete(this.freqList, node.freq-1)
		this.freqList[node.freq] = ll
		return
	}

	// delete useless ll
	if ll.len == 1 {
		ll.prev.next = ll.next
		ll.next.prev = ll.prev
		delete(this.freqList, node.freq-1)
	}

	// make new ll, and insert to ll's next
	if this.freqList[node.freq] == nil {
		this.freqList[node.freq] = newDoublyLL()
		this.freqList[node.freq].next = ll.next
		ll.next.prev = this.freqList[node.freq]
		this.freqList[node.freq].prev = ll
		ll.next = this.freqList[node.freq]
	}
	ll.remove(node)
	this.freqList[node.freq].pushToTail(node)
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	node := this.nodeAddr[key]
	ll := this.freqList[node.freq]
	node.freq--

	// delete node
	if node.freq == 0 {
		delete(this.nodeAddr, key)
		ll.remove(node)
		// delete this ll
		if ll.len == 0 {
			ll.prev.next = ll.next
			ll.next.prev = ll.prev
			delete(this.freqList, node.freq+1)
		}
		return
	}

	// just change ll's freq
	if this.freqList[node.freq] == nil && ll.len == 1 {
		delete(this.freqList, node.freq+1)
		this.freqList[node.freq] = ll
		return
	}

	// delete useless ll
	if ll.len == 1 {
		ll.prev.next = ll.next
		ll.next.prev = ll.prev
		delete(this.freqList, node.freq+1)
	}

	// make new ll, and insert to ll's prev
	if this.freqList[node.freq] == nil {
		this.freqList[node.freq] = newDoublyLL()
		this.freqList[node.freq].prev = ll.prev
		ll.prev.next = this.freqList[node.freq]
		this.freqList[node.freq].next = ll
		ll.prev = this.freqList[node.freq]
	}
	ll.remove(node)
	this.freqList[node.freq].pushToTail(node)
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.tail.prev.len > 0 {
		return this.tail.prev.peekTail().key
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.head.next.len > 0 {
		return this.head.next.peekTail().key
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

type doublyLL struct {
	prev, next *doublyLL
	head, tail *dataNode
	len        int
}

func newDoublyLL() *doublyLL {
	head := newDataNode("")
	tail := newDataNode("")
	head.next = tail
	tail.prev = head
	return &doublyLL{head: head, tail: tail}
}

func (ll *doublyLL) pushToTail(node *dataNode) {
	node.prev = ll.tail.prev
	node.next = ll.tail
	node.prev.next = node
	node.next.prev = node
	ll.len++
}

func (ll *doublyLL) remove(node *dataNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	ll.len--
}

func (ll doublyLL) peekTail() *dataNode {
	return ll.tail.prev
}

type dataNode struct {
	key        string
	freq       int
	prev, next *dataNode
}

func newDataNode(key string) *dataNode {
	return &dataNode{key: key}
}
