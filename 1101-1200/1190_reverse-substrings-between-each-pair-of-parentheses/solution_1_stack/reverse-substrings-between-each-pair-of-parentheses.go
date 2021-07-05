package main

// Tags:
// Stack

func reverseParentheses(s string) string {
	stk := RuneStack{}
	for _, ch := range s {
		if ch == ')' {
			tmp := []rune{}
			for stk.Top() != '(' {
				tmp = append(tmp, stk.Pop())
			}
			stk.Pop()
			stk.Push(tmp...)
		} else {
			stk.Push(ch)
		}
	}
	return string(stk)
}

type RuneStack []rune

func (s RuneStack) Len() int        { return len(s) }
func (s RuneStack) Top() rune       { return s[s.Len()-1] }
func (s *RuneStack) Push(x ...rune) { *s = append(*s, x...) }
func (s *RuneStack) Pop() rune {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}
