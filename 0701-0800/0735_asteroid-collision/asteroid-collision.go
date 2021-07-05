package main

// Tags:
// Stack

func asteroidCollision(asteroids []int) []int {
	stk := IntStack{}
	for _, a := range asteroids {
		if a > 0 {
			stk.Push(a)
			continue
		}
		ok := true
		for stk.Len() > 0 && stk.Top() > 0 {
			top := stk.Top()
			if top <= -a {
				stk.Pop()
			}
			if top >= -a {
				ok = false
				break
			}
		}
		if ok {
			stk.Push(a)
		}
	}
	return stk
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
