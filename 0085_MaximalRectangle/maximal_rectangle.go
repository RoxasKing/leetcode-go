package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	out := 0
	for i := 0; i < len(matrix); i++ {
		heights := make([]int, 0)
		// 计算当前行起每一列只包含 1 的最大高度
		for j := 0; j < len(matrix[0]); j++ {
			k := i
			for k < len(matrix) && matrix[k][j] == '1' {
				k++
			}
			heights = append(heights, k-i)
		}
		out = max(out, largestArea(heights))
	}
	return out
}

func largestArea(heights []int) int {
	// 定义一个切片，当作栈用
	stack := make([]int, 0)
	// 初始化栈底为 -1
	stack = append(stack, -1)
	// 记录最大面积
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		// 若当前遍历到的数字小于等于栈顶数字
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= heights[i] {
			maxArea = max(maxArea, heights[stack[len(stack)-1]]*(i-stack[len(stack)-2]-1))
			// 栈顶元素出栈
			stack = stack[:len(stack)-1]
		}
		// 当前下标入栈
		stack = append(stack, i)
	}
	for stack[len(stack)-1] != -1 {
		maxArea = max(maxArea, heights[stack[len(stack)-1]]*(len(heights)-stack[len(stack)-2]-1))
		stack = stack[:len(stack)-1]
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalRectangle(matrix))
}
