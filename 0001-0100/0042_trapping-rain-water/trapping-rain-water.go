package main

// Tags:
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

// Two Pointers
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
