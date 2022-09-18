package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	lMax, rMax := make([]int, n), make([]int, n)
	lMax[0], rMax[n-1] = height[0], height[n-1]
	for i := 1; i < n; i++ {
		lMax[i] = max(lMax[i-1], height[i])
		rMax[n-1-i] = max(rMax[n-i], height[n-1-i])
	}
	var out int
	for i := range height {
		out += min(lMax[i], rMax[i]) - height[i]
	}
	return out
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
