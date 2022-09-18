package main

// Difficulty:
// Hard

// Tags:
// Two Pointers

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	l, r := 0, n-1
	lMax, rMax := -1<<31, -1<<31
	var out int
	for l < r {
		lMax = max(lMax, height[l])
		rMax = max(rMax, height[r])
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
