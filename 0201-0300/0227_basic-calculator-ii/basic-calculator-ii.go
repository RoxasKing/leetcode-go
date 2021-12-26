package main

// Difficulty:
// Medium

// Tags:
// Stack

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

func calculate(s string) int {
	stk := IntStack{}
	num, op := 0, '+'
	for _, ch := range s + "+" {
		if ch == ' ' {
			continue
		}
		if '0' <= ch && ch <= '9' {
			num = num*10 + int(ch-'0')
			continue
		}
		switch op {
		case '+':
			stk.Push(num)
		case '-':
			stk.Push(-num)
		case '*':
			stk.Push(stk.Pop() * num)
		case '/':
			stk.Push(stk.Pop() / num)
		}
		num, op = 0, ch
	}
	out := 0
	for _, num := range stk {
		out += num
	}
	return out
}
