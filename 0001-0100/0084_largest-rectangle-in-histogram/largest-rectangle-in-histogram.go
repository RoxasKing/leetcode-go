package main

// Tags:
// Monotone Stack
func largestRectangleArea(heights []int) int {
	out := 0
	s := IntStack{-1}
	for i := range heights {
		for s.Len() > 1 && heights[i] <= heights[s.Top()] {
			out = Max(out, heights[s.Pop()]*(i-1-s.Top()))
		}
		s.Push(i)
	}
	for s.Len() > 1 {
		out = Max(out, heights[s.Pop()]*(len(heights)-1-s.Top()))
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

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
