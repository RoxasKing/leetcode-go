package main

// Difficulty:
// Medium

// Tags:
// Linked List

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head, tail *Node
	cnt        int
}

func Constructor() MyLinkedList {
	head := &Node{-1, nil}
	head.next = head
	tail := head
	return MyLinkedList{head, tail, 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.cnt {
		return -1
	}
	ptr := this.head.next
	for i := 0; i < index; i++ {
		ptr = ptr.next
	}
	return ptr.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{val, this.head.next}
	this.head.next = node
	if this.head == this.tail {
		this.tail = node
	}
	this.cnt++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{val, this.tail.next}
	this.tail.next = node
	this.tail = node
	this.cnt++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.cnt {
		return
	}
	if index == this.cnt {
		this.AddAtTail(val)
		return
	}
	ptr := this.head
	for i := 0; i < index; i++ {
		ptr = ptr.next
	}
	ptr.next = &Node{val, ptr.next}
	this.cnt++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.cnt {
		return
	}
	ptr := this.head
	for i := 0; i < index; i++ {
		ptr = ptr.next
	}
	if ptr.next == this.tail {
		this.tail = ptr
	}
	ptr.next = ptr.next.next
	this.cnt--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
