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
	stack []int
	help  []int
}

/** initialize your data structure here. */
func NewMinStack() MinStack {
	return MinStack{
		stack: []int{},
		help:  []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.help) > 0 && this.help[len(this.help)-1] < x {
		return
	}
	this.help = append(this.help, x)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if this.help[len(this.help)-1] == val {
		this.help = this.help[:len(this.help)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	if len(this.help) == 0 {
		return -1
	}
	return this.help[len(this.help)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
