package main

/*
  Implement a last in first out (LIFO) stack using only two queues. The implemented stack should support all the functions of a normal queue (push, top, pop, and empty).

  Implement the MyStack class:
    1. void push(int x) Pushes element x to the top of the stack.
    2. int pop() Removes the element on the top of the stack and returns it.
    3. int top() Returns the element on the top of the stack.
    4. boolean empty() Returns true if the stack is empty, false otherwise.

  Notes:
    1. You must use only standard operations of a queue, which means only push to back, peek/pop from front, size, and is empty operations are valid.
    2. Depending on your language, the queue may not be supported natively. You may simulate a queue using a list or deque (double-ended queue), as long as you use only a queue's standard operations.


  Example 1:
    Input
      ["MyStack", "push", "push", "top", "pop", "empty"]
      [[], [1], [2], [], [], []]
    Output
      [null, null, null, 2, 2, false]
    Explanation
      MyStack myStack = new MyStack();
      myStack.push(1);
      myStack.push(2);
      myStack.top(); // return 2
      myStack.pop(); // return 2
      myStack.empty(); // return False

  Constraints:
    1. 1 <= x <= 9
    2. At most 100 calls will be made to push, pop, top, and empty.
    3. All the calls to pop and top are valid.

    Follow-up: Can you implement the stack such that each operation is amortized O(1) time complexity? In other words, performing n operations will take overall O(n) time even if one of those operations may take longer. You can use more than two queues.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/implement-stack-using-queues
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
