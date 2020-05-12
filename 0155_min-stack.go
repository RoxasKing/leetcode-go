package leetcode

// MinStack ...
type MinStack struct {
	stack []int
	help  []int
}

// NewMinStack ...
func NewMinStack() MinStack {
	return MinStack{stack: []int{}, help: []int{}}
}

// Push ...
func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	if len(s.help) == 0 || x <= s.help[len(s.help)-1] {
		s.help = append(s.help, x)
	}
}

// Pop ...
func (s *MinStack) Pop() {
	if s.stack[len(s.stack)-1] == s.help[len(s.help)-1] {
		s.help = s.help[:len(s.help)-1]
	}
	s.stack = s.stack[:len(s.stack)-1]
}

// Top ...
func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

// GetMin ...
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
