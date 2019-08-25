package main

import "fmt"

func largestRectangleArea(heights []int) int {
	// 定义一个切片，当作栈用
	stack := make([]int, 0)
	// 初始化栈底为 -1
	stack = append(stack, -1)
	// 记录最大面积
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		// 若当前遍历到的数字小于等于栈顶数字
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= heights[i] {
			maxArea = max(maxArea, heights[stack[len(stack)-1]]*(i-1-stack[len(stack)-2]))
			// 栈顶元素出栈
			stack = stack[:len(stack)-1]
		}
		// 当前下标入栈
		stack = append(stack, i)
	}
	for stack[len(stack)-1] != -1 {
		maxArea = max(maxArea, heights[stack[len(stack)-1]]*(len(heights)-1-stack[len(stack)-2]))
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
	height := []int{2, 1, 5, 6, 2, 3}
	//height = []int{5, 5, 5, 5, 5}
	fmt.Println(largestRectangleArea(height))
}
