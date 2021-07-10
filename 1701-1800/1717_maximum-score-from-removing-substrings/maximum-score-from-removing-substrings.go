package main

// Tags:
// Stack

func maximumGain(s string, x int, y int) int {
	a, b := byte('a'), byte('b')
	if x < y {
		x, y = y, x
		a, b = b, a
	}
	out := 0
	stk1 := ByteStack{}
	for i := range s {
		if s[i] == b && stk1.Len() > 0 && stk1.Top() == a {
			stk1.Pop()
			out += x
		} else {
			stk1.Push(s[i])
		}
	}
	stk2 := ByteStack{}
	for _, c := range stk1 {
		if c == a && stk2.Len() > 0 && stk2.Top() == b {
			stk2.Pop()
			out += y
		} else {
			stk2.Push(c)
		}
	}
	return out
}

type ByteStack []byte

func (s ByteStack) Len() int     { return len(s) }
func (s ByteStack) Top() byte    { return s[s.Len()-1] }
func (s *ByteStack) Push(x byte) { *s = append(*s, x) }
func (s *ByteStack) Pop() byte {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}
