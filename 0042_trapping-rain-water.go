package leetcode

/*
  给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/

// Double Pointer
func trap(height []int) int {
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

// Stack
func trap2(height []int) int {
	var out int
	var stack []int
	for i := 0; i < len(height); i++ {
		for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
			bottomHeight := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			distance := i - 1 - stack[len(stack)-1]
			boundedHeight := Min(height[i], height[stack[len(stack)-1]]) - bottomHeight
			out += distance * boundedHeight
		}
		stack = append(stack, i)
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
