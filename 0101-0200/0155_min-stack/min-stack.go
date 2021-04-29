package main

/*
  Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

  Implement the MinStack class:
    1. MinStack() initializes the stack object.
    2. void push(val) pushes the element val onto the stack.
    3. void pop() removes the element on the top of the stack.
    4. int top() gets the top element of the stack.
    5. int getMin() retrieves the minimum element in the stack.

  Example 1:
    Input
      ["MinStack","push","push","push","getMin","pop","top","getMin"]
      [[],[-2],[0],[-3],[],[],[],[]]
    Output
      [null,null,null,null,-3,null,0,-2]
    Explanation
      MinStack minStack = new MinStack();
      minStack.push(-2);
      minStack.push(0);
      minStack.push(-3);
      minStack.getMin(); // return -3
      minStack.pop();
      minStack.top();    // return 0
      minStack.getMin(); // return -2

  Constraints:
    1. -2^31 <= val <= 2^31 - 1
    2. Methods pop, top and getMin operations will always be called on non-empty stacks.
    3. At most 3 * 10^4 calls will be made to push, pop, top, and getMin.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/min-stack
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type MinStack struct {
	nums []int
	mins []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.nums = append(this.nums, val)
	if len(this.mins) == 0 || this.mins[len(this.mins)-1] >= val {
		this.mins = append(this.mins, val)
	}
}

func (this *MinStack) Pop() {
	if this.mins[len(this.mins)-1] == this.nums[len(this.nums)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}
	this.nums = this.nums[:len(this.nums)-1]
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
