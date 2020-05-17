package codinginterviews

/*
  用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type CQueue struct {
	stackA *Stack
	stackB *Stack
}

func NewCQueue() CQueue {
	return CQueue{
		stackA: NewStack(),
		stackB: NewStack(),
	}
}

func (c *CQueue) AppendTail(value int) {
	for c.stackB.Size() != 0 {
		value := c.stackB.Pop()
		c.stackA.Push(value)
	}
	c.stackA.Push(value)
}

func (c *CQueue) DeleteHead() int {
	for c.stackA.Size() != 0 {
		value := c.stackA.Pop()
		c.stackB.Push(value)
	}
	return c.stackB.Pop()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
