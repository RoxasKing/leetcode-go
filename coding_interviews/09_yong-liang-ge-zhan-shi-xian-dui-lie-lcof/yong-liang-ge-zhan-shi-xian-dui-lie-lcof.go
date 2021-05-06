package main

/*
  用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

  示例 1：
    输入：
    ["CQueue","appendTail","deleteHead","deleteHead"]
    [[],[3],[],[]]
    输出：[null,null,3,-1]

  示例 2：
    输入：
    ["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
    [[],[],[5],[2],[],[]]
    输出：[null,-1,null,null,5,2]

  提示：
    1 <= values <= 10000
    最多会对 appendTail、deleteHead 进行 10000 次调用

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type CQueue struct {
	a *IntStack
	b *IntStack
}

func Constructor() CQueue {
	return CQueue{
		a: &IntStack{},
		b: &IntStack{},
	}
}

func (this *CQueue) AppendTail(value int) {
	for this.b.Len() > 0 {
		this.a.Push(this.b.Pop())
	}
	this.a.Push(value)
}

func (this *CQueue) DeleteHead() int {
	for this.a.Len() > 0 {
		this.b.Push(this.a.Pop())
	}
	if this.b.Len() == 0 {
		return -1
	}
	return this.b.Pop()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

type IntStack []int

func (s IntStack) Len() int    { return len(s) }
func (s IntStack) Top() int    { return s[s.Len()-1] }
func (s *IntStack) Push(x int) { *s = append(*s, x) }
func (s *IntStack) Pop() int {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}
