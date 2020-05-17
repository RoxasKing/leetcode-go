package codinginterviews

import "testing"

func TestMinStack(t *testing.T) {
	minStack := NewMinStack()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	if minStack.Min() != -3 {
		t.Errorf("minStack.Min() = %v, want %v", minStack.Min(), -3)
	}
	minStack.Pop()
	if minStack.Top() != 0 {
		t.Errorf("minStack.Top() = %v, want %v", minStack.Top(), 0)
	}
	if minStack.Min() != -2 {
		t.Errorf("minStack.Min() = %v, want %v", minStack.Min(), -2)
	}
}
