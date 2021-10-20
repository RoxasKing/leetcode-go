package main

// Tags:
// Monotonic Stack
func removeKdigits(num string, k int) string {
	stk := CharStack{}

	for i := range num {
		for ; stk.Len() > 0 && stk.Top() > num[i] && k > 0; k-- {
			stk.Pop()
		}
		stk.Push(num[i])
	}

	for ; k > 0; k-- {
		stk.Pop()
	}

	for stk.Len() > 0 && stk[0] == '0' {
		stk = stk[1:]
	}

	if len(stk) == 0 {
		return "0"
	}
	return string(stk)
}

type CharStack []byte

func (s CharStack) Len() int     { return len(s) }
func (s CharStack) Top() byte    { return s[s.Len()-1] }
func (s *CharStack) Push(x byte) { *s = append(*s, x) }
func (s *CharStack) Pop() byte {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}
