package main

// Tags:
// Monotone Stack
func trap(height []int) int {
	stack := []int{}
	out := 0
	for i := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			bottom := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			out += (i - 1 - stack[len(stack)-1]) * (Min(height[stack[len(stack)-1]], height[i]) - bottom)
		}
		stack = append(stack, i)
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
