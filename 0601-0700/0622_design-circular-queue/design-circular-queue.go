package main

// Difficulty:
// Medium

// Tags:
// Design

type MyCircularQueue struct {
	a       []int
	k, l, r int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{make([]int, k+1), k + 1, 0, 0}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.r = (this.r + 1) % this.k
	this.a[this.r] = value
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.l = (this.l + 1) % this.k
	return true
}

func (this MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.a[(this.l+1)%this.k]
}

func (this MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.a[this.r]
}

func (this MyCircularQueue) IsEmpty() bool {
	return this.l == this.r
}

func (this MyCircularQueue) IsFull() bool {
	return (this.r+1)%this.k == this.l
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
