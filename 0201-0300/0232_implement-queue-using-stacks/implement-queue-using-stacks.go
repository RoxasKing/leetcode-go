package main

type MyQueue struct {
	a IntStack
	b IntStack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	for this.b.Len() > 0 {
		this.a.Push(this.b.Pop())
	}
	this.a.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	for this.a.Len() > 0 {
		this.b.Push(this.a.Pop())
	}
	return this.b.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	for this.a.Len() > 0 {
		this.b.Push(this.a.Pop())
	}
	return this.b.Top()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.a.Len() == 0 && this.b.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type IntStack []int

func (s IntStack) Len() int { return len(s) }
func (s IntStack) Top() int { return s[len(s)-1] }
func (s *IntStack) Pop() int {
	last := len(*s) - 1
	out := (*s)[last]
	*s = (*s)[:last]
	return out
}
func (s *IntStack) Push(x int) { *s = append(*s, x) }
