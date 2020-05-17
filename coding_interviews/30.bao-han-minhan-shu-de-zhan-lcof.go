package codinginterviews

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
	stack *Stack
	help  *Stack
}

/** initialize your data structure here. */
func NewMinStack() MinStack {
	return MinStack{
		stack: NewStack(),
		help:  NewStack(),
	}
}

func (this *MinStack) Push(x int) {
	this.stack.Push(x)
	if this.help.Size() == 0 || this.help.Peek() >= x {
		this.help.Push(x)
	}
}

func (this *MinStack) Pop() {
	if this.stack.Size() == 0 {
		return
	}
	val := this.stack.Pop()
	if this.help.Peek() == val {
		this.help.Pop()
	}
}

func (this *MinStack) Top() int {
	return this.stack.Peek()
}

func (this *MinStack) Min() int {
	return this.help.Peek()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
