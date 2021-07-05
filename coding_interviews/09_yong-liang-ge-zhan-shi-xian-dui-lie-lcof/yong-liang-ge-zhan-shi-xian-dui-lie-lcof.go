package main

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
