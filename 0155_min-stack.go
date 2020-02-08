package leetcode

// MinStack ...
type MinStack struct {
	stackArray []int
	helpArray  []int
}

// NewMinStack ...
func NewMinStack() MinStack {
	return MinStack{stackArray: make([]int, 0), helpArray: make([]int, 0)}
}

// Push ...
func (s *MinStack) Push(x int) {
	s.stackArray = append(s.stackArray, x)
	if len(s.helpArray) == 0 ||
		x <= s.helpArray[len(s.helpArray)-1] {
		s.helpArray = append(s.helpArray, x)
	}
}

// Pop ...
func (s *MinStack) Pop() {
	if s.stackArray[len(s.stackArray)-1] ==
		s.helpArray[len(s.helpArray)-1] {
		s.helpArray = s.helpArray[:len(s.helpArray)-1]
	}
	s.stackArray = s.stackArray[:len(s.stackArray)-1]

}

// Top ...
func (s *MinStack) Top() int {
	return s.stackArray[len(s.stackArray)-1]
}

// GetMin ...
func (s *MinStack) GetMin() int {
	return s.helpArray[len(s.helpArray)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
