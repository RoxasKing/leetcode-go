package main

/*
  给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
  求在该柱状图中，能够勾勒出来的矩形的最大面积。

  以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。

  图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。

  示例:
    输入: [2,1,5,6,2,3]
    输出: 10

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/largest-rectangle-in-histogram
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Stack
func largestRectangleArea(heights []int) int {
	out := 0
	stack := MakeStack()
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

type Stack []int

func MakeStack() Stack {
	return Stack{}
}

func (s Stack) Size() int {
	return len(s)
}

func (s Stack) Top() int {
	last := len(s) - 1
	return s[last]
}

func (s *Stack) Pop() int {
	last := len(*s) - 1
	out := (*s)[last]
	*s = (*s)[:last]
	return out
}

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
