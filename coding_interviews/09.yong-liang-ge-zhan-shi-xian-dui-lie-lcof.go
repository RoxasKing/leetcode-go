package codinginterviews

/*
  用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// CQueue ...
type CQueue struct {
	stackA *stack
	stackB *stack
}

// Constructor ...
func Constructor() CQueue {
	return CQueue{
		stackA: newStack(),
		stackB: newStack(),
	}
}

// AppendTail ...
func (c *CQueue) AppendTail(value int) {
	for c.stackB.size() != 0 {
		value := c.stackB.pop()
		c.stackA.push(value)
	}
	c.stackA.push(value)
}

// DeleteHead ...
func (c *CQueue) DeleteHead() int {
	for c.stackA.size() != 0 {
		value := c.stackA.pop()
		c.stackB.push(value)
	}
	return c.stackB.pop()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

type stack struct {
	elems []int
}

func newStack() *stack {
	return &stack{
		elems: []int{},
	}
}

func (s *stack) size() int {
	return len(s.elems)
}

func (s *stack) push(value int) {
	s.elems = append(s.elems, value)
}

func (s *stack) pop() int {
	if s.size() == 0 {
		return -1
	}
	value := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return value
}
