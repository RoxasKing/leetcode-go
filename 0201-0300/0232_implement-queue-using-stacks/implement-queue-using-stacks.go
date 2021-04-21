package main

/*
  Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

  Implement the MyQueue class:

  void push(int x) Pushes element x to the back of the queue.
  int pop() Removes the element from the front of the queue and returns it.
  int peek() Returns the element at the front of the queue.
  boolean empty() Returns true if the queue is empty, false otherwise.
  Notes:

  You must use only standard operations of a stack, which means only push to top, peek/pop from top, size, and is empty operations are valid.
  Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.
  Follow-up: Can you implement the queue such that each operation is amortized O(1) time complexity? In other words, performing n operations will take overall O(n) time even if one of those operations may take longer.

  Example 1:
    Input
      ["MyQueue", "push", "push", "peek", "pop", "empty"]
      [[], [1], [2], [], [], []]
    Output
      [null, null, null, 1, 1, false]
    Explanation
      MyQueue myQueue = new MyQueue();
      myQueue.push(1); // queue is: [1]
      myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
      myQueue.peek(); // return 1
      myQueue.pop(); // return 1, queue is [2]
      myQueue.empty(); // return false

  Constraints:
    1. 1 <= x <= 9
    2. At most 100 calls will be made to push, pop, peek, and empty.
    3. All the calls to pop and peek are valid.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/implement-queue-using-stacks
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
