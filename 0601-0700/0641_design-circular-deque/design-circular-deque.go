package main

// Difficulty:
// Medium

// Tags:
// Linked List

type listNode struct {
	val        int
	prev, next *listNode
}

type MyCircularDeque struct {
	remain     int
	head, tail *listNode
}

func Constructor(k int) MyCircularDeque {
	head, tail := &listNode{}, &listNode{}
	head.next = tail
	tail.prev = head
	return MyCircularDeque{k, head, tail}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.remain--
	node := &listNode{val: value}
	this.head.next.prev = node
	node.next = this.head.next
	node.prev = this.head
	this.head.next = node
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.remain--
	node := &listNode{val: value}
	this.tail.prev.next = node
	node.prev = this.tail.prev
	node.next = this.tail
	this.tail.prev = node
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.remain++
	this.head.next.next.prev = this.head
	this.head.next = this.head.next.next
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.remain++
	this.tail.prev.prev.next = this.tail
	this.tail.prev = this.tail.prev.prev
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.head.next.val
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.tail.prev.val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.head.next == this.tail
}

func (this *MyCircularDeque) IsFull() bool {
	return this.remain == 0
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
