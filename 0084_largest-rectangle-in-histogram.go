package main

/*
  给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
  求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/

func largestRectangleArea(heights []int) int {
	stack := append(make([]int, 0, len(heights)+1), -1)
	var out int
	for i := range heights {
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= heights[i] {
			out = Max(out, heights[stack[len(stack)-1]]*(i-1-stack[len(stack)-2]))
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for stack[len(stack)-1] != -1 {
		out = Max(out, heights[stack[len(stack)-1]]*(len(heights)-1-stack[len(stack)-2]))
		stack = stack[:len(stack)-1]
	}
	return out
}
