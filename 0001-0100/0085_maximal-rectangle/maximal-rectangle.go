package main

// Difficulty:
// Hard

// Tags:
// Monotonic Stack

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	heights := make([]int, n)
	out := 0
	cal := func() {
		max := 0
		stk := IntStack{-1}
		for i, height := range heights {
			for stk.Top() != -1 && heights[stk.Top()] >= height {
				max = Max(max, heights[stk.Pop()]*(i-1-stk.Top()))
			}
			stk.Push(i)
		}
		for stk.Top() != -1 {
			max = Max(max, heights[stk.Pop()]*(n-1-stk.Top()))
		}
		out = Max(out, max)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		cal()
	}
	return out
}

type IntStack []int

func (s IntStack) Len() int      { return len(s) }
func (s IntStack) Top() int      { return s[s.Len()-1] }
func (s *IntStack) Push(num int) { *s = append(*s, num) }
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
