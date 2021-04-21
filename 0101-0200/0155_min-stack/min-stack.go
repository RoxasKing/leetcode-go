package main

/*
  设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
    push(x) —— 将元素 x 推入栈中。
    pop() —— 删除栈顶的元素。
    top() —— 获取栈顶元素。
    getMin() —— 检索栈中的最小元素。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/min-stack
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type MinStack struct {
	stack []int
	help  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	if len(s.help) == 0 || x <= s.help[len(s.help)-1] {
		s.help = append(s.help, x)
	}
}

func (s *MinStack) Pop() {
	if s.stack[len(s.stack)-1] == s.help[len(s.help)-1] {
		s.help = s.help[:len(s.help)-1]
	}
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.help[len(s.help)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
