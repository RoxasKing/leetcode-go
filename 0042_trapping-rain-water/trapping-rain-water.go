package main

/*
  给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/

// Stack
func trap(height []int) int {
	var out int
	stack := MakeStack()
	for i := 0; i < len(height); i++ {
		for stack.Size() > 0 && height[stack.Top()] < height[i] {
			baseHeight := height[stack.Pop()]
			if stack.Size() == 0 {
				break
			}
			h := Min(height[stack.Top()], height[i]) - baseHeight
			w := i - 1 - stack.Top()
			out += h * w
		}
		stack.Push(i)
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

// Double Pointer
func trap2(height []int) int {
	if len(height) == 0 {
		return 0
	}
	l, r := 0, len(height)-1
	lmax, rmax := height[l], height[r]
	var out int
	for l <= r {
		lmax = Max(lmax, height[l])
		rmax = Max(rmax, height[r])
		if lmax < rmax {
			out += lmax - height[l]
			l++
		} else {
			out += rmax - height[r]
			r--
		}
	}
	return out
}

// Dynamic Programming
func trap3(height []int) int {
	if len(height) == 0 {
		return 0
	}
	lmax, rmax := make([]int, len(height)), make([]int, len(height))
	lmax[0], rmax[len(height)-1] = height[0], height[len(height)-1]
	for i := 1; i < len(height); i++ {
		lmax[i] = Max(lmax[i-1], height[i])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rmax[i] = Max(rmax[i+1], height[i])
	}
	var out int
	for i := range height {
		out += Min(lmax[i], rmax[i]) - height[i]
	}
	return out
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
