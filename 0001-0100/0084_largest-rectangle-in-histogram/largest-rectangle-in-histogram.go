package main

// Difficulty:
// Hard

// Tags:
// Monotonic Stack

func largestRectangleArea(heights []int) int {
	out := 0
	stk := []int{-1}
	for i := range heights {
		for len(stk) > 1 && heights[stk[len(stk)-1]] > heights[i] {
			height := heights[stk[len(stk)-1]]
			stk = stk[:len(stk)-1]
			out = Max(out, height*(i-1-stk[len(stk)-1]))
		}
		stk = append(stk, i)
	}
	for len(stk) > 1 {
		height := heights[stk[len(stk)-1]]
		stk = stk[:len(stk)-1]
		out = Max(out, height*(len(heights)-1-stk[len(stk)-1]))
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
