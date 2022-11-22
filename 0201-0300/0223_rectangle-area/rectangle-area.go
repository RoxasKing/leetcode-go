package main

// Difficulty:
// Medium

// Tags:
// Math

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	o := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1)
	l, r := max(ax1, bx1), min(ax2, bx2)
	u, d := min(ay2, by2), max(ay1, by1)
	if l < r && u > d {
		o -= (r - l) * (u - d)
	}
	return o
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
