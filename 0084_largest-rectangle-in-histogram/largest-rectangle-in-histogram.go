package main

/*
  给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
  求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/

// Stack
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
