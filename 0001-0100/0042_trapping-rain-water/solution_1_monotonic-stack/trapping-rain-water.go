package main

// Difficulty:
// Hard

// Tags:
// Monotonic Stack

func trap(height []int) int {
	stk := []int{}
	top := func() int { return len(stk) - 1 }
	o := 0
	for i, x := range height {
		for len(stk) > 0 && height[stk[top()]] < x {
			lo := height[stk[top()]]
			if stk = stk[:top()]; len(stk) == 0 {
				break
			}
			hi := min(x, height[stk[top()]])
			o += (i - 1 - stk[top()]) * (hi - lo)
		}
		stk = append(stk, i)
	}
	return o
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
