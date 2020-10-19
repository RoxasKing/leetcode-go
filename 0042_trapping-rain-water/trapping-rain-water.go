package main

/*
  给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

  上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

  示例:
    输入: [0,1,0,2,1,0,1,3,2,1,2,1]
    输出: 6

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/trapping-rain-water
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Stack
func trap(height []int) int {
	var out int
	stack := MakeStack()
	for i := range height {
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
	(*s) = (*s)[:last]
	return out
}

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

// Double Pointer
func trap2(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	l, r := 0, n-1
	lMax, rMax := -1<<31, -1<<31
	var out int
	for l < r {
		lMax = Max(lMax, height[l])
		rMax = Max(rMax, height[r])
		if lMax < rMax {
			out += lMax - height[l]
			l++
		} else {
			out += rMax - height[r]
			r--
		}
	}
	return out
}

// Dynamic Programming
func trap3(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	lMax, rMax := make([]int, n), make([]int, n)
	lMax[0], rMax[n-1] = height[0], height[n-1]
	for i := 1; i < n; i++ {
		lMax[i] = Max(lMax[i-1], height[i])
		rMax[n-1-i] = Max(rMax[n-i], height[n-1-i])
	}
	var out int
	for i := range height {
		out += Min(lMax[i], rMax[i]) - height[i]
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
