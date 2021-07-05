package main

import "strconv"

// Stack

func calPoints(ops []string) int {
	s := IntStack{}
	for _, op := range ops {
		switch op {
		case "+":
			top := s.Len() - 1
			s.Push(s[top] + s[top-1])
		case "D":
			s.Push(s.Top() << 1)
		case "C":
			s.Pop()
		default:
			num, _ := strconv.Atoi(op)
			s.Push(num)
		}
	}

	out := 0
	for _, num := range s {
		out += num
	}
	return out
}

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
