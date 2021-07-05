package main

// Tags:
// Stack
/* preprocess parentheses */

func reverseParentheses(s string) string {
	n := len(s)
	pair := make([]int, n)
	stk := IntStack{}
	for i, b := range s {
		if b == '(' {
			stk.Push(i)
		} else if b == ')' {
			j := stk.Pop()
			pair[i], pair[j] = j, i
		}
	}

	out := []byte{}
	for i, step := 0, 1; i < n; i += step {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			step = -step
		} else {
			out = append(out, s[i])
		}
	}
	return string(out)
}

type IntStack []int

func (s IntStack) Len() int       { return len(s) }
func (s IntStack) Top() int       { return s[s.Len()-1] }
func (s *IntStack) Push(x ...int) { *s = append(*s, x...) }
func (s *IntStack) Pop() int {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}
