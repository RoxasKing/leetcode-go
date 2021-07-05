package main

// Tags:
// Stack

func lengthLongestPath(input string) int {
	n := len(input)
	stk1 := IntStack{} // storage parent dir path length
	stk2 := IntStack{} // storage full dir path length

	out := 0
	for l, r := 0, 0; r < n; l, r = r+1, r+1 {
		for ; r < n && input[r] == '\t'; r++ {
		}
		for stk1.Len() > r-l {
			stk1.Pop()
			stk2.Pop()
		}

		isFile := false
		for l = r; r < n && input[r] != '\n'; r++ {
			isFile = isFile || input[r] == '.'
		}

		if stk1.Len() > 0 {
			stk1.Push(r - l + 1)
		} else {
			stk1.Push(r - l) // root path, not need to add '/'
		}
		if stk2.Len() > 0 {
			stk2.Push(stk2.Top() + stk1.Top()) // acumulate path length
		} else {
			stk2.Push(stk1.Top())
		}

		if isFile && stk2.Top() > out {
			out = stk2.Top()
		}
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
