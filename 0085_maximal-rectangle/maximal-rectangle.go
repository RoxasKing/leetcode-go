package main

/*
  给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
*/

// Stack
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var out int
	heights := make([]int, len(matrix[0]))
	for i := range matrix {
		for j := range matrix[0] {
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
	stack := MakeStack()
	stack.Push(-1)
	var out int
	for i := range heights {
		for stack.Top() != -1 && heights[i] <= heights[stack.Top()] {
			out = Max(out, heights[stack.Pop()]*(i-1-stack.Top()))
		}
		stack.Push(i)
		stack = append(stack, i)
	}
	for stack[len(stack)-1] != -1 {
		out = Max(out, heights[stack.Pop()]*(len(heights)-1-stack.Top()))
	}
	return out
}

type Stack []int

func MakeStack() Stack {
	return Stack{}
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) Pop() int {
	last := len(*s) - 1
	res := (*s)[last]
	(*s) = (*s)[:last]
	return res
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
