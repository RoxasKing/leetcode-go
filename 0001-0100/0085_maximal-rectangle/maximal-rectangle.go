package main

// Tags:
// Stack
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var out int
	m, n := len(matrix), len(matrix[0])
	heights := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		out = Max(out, largestRectangleArea(heights))
	}
	return out
}

func largestRectangleArea(heights []int) int {
	out := 0
	stack := MakeIntStack()
	stack.Push(-1)
	for i := range heights {
		for stack.Top() != -1 && heights[i] <= heights[stack.Top()] {
			out = Max(out, heights[stack.Pop()]*(i-1-stack.Top()))
		}
		stack.Push(i)
	}
	for stack.Top() != -1 {
		out = Max(out, heights[stack.Pop()]*(len(heights)-1-stack.Top()))
	}
	return out
}

type IntStack []int

func MakeIntStack() IntStack {
	return IntStack{}
}

func (s IntStack) Size() int {
	return len(s)
}

func (s IntStack) Top() int {
	last := len(s) - 1
	return s[last]
}

func (s *IntStack) Pop() int {
	last := len(*s) - 1
	out := (*s)[last]
	*s = (*s)[:last]
	return out
}

func (s *IntStack) Push(num int) {
	*s = append(*s, num)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
