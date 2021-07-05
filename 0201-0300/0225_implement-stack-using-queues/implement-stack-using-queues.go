package main

type MyStack struct {
	q1, q2 *IntQueue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		q1: &IntQueue{},
		q2: &IntQueue{},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	for this.q2.Len() > 0 {
		this.q1.PushBack(this.q2.Shift())
	}
	this.q1.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for this.q2.Len() > 0 {
		this.q1.PushBack(this.q2.Shift())
	}
	for this.q1.Len() > 1 {
		this.q2.PushBack(this.q1.Shift())
	}
	return this.q1.Shift()
}

/** Get the top element. */
func (this *MyStack) Top() int {
	for this.q2.Len() > 0 {
		this.q1.PushBack(this.q2.Shift())
	}
	for this.q1.Len() > 1 {
		this.q2.PushBack(this.q1.Shift())
	}
	out := this.q1.Head()
	this.q2.PushBack(this.q1.Shift())
	return out
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.q1.Len() == 0 && this.q2.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

type IntQueue []int

func (q IntQueue) Len() int        { return len(q) }
func (q IntQueue) Head() int       { return q[0] }
func (q *IntQueue) PushBack(x int) { *q = append(*q, x) }
func (q *IntQueue) Shift() int {
	out := (*q)[0]
	*q = (*q)[1:]
	return out
}
