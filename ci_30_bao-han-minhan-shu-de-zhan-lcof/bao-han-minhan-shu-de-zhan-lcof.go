package main

/*
  定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

  示例:
    MinStack minStack = new MinStack();
    minStack.push(-2);
    minStack.push(0);
    minStack.push(-3);
    minStack.min();      --> 返回 -3.
    minStack.pop();
    minStack.top();      --> 返回 0.
    minStack.min();      --> 返回 -2.

  提示：
  各函数的调用总次数不超过 20000 次

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type MinStack struct {
	stk []int
	min []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stk = append(this.stk, x)
	if len(this.min) == 0 || x <= this.min[len(this.min)-1] {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	if this.min[len(this.min)-1] == this.stk[len(this.stk)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.stk = this.stk[:len(this.stk)-1]
}

func (this *MinStack) Top() int {
	return this.stk[len(this.stk)-1]
}

func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
