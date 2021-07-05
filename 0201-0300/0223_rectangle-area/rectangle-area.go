package main

// Tags:
// Math

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	out := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1)
	l, r := Max(ax1, bx1), Min(ax2, bx2)
	u, d := Min(ay2, by2), Max(ay1, by1)
	if l < r && u > d {
		out -= (r - l) * (u - d)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
