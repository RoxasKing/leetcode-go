package main

// Difficulty:
// Hard

// Tags:
// Math
// Hash

func isRectangleCover(rectangles [][]int) bool {
	freq := map[[2]int]int{}
	area := 0
	mx, my, ma, mb := 1<<31-1, 1<<31-1, -1<<31, -1<<31
	for _, rec := range rectangles {
		x, y, a, b := rec[0], rec[1], rec[2], rec[3]
		freq[[2]int{x, y}]++
		freq[[2]int{x, b}]++
		freq[[2]int{a, y}]++
		freq[[2]int{a, b}]++
		area += (a - x) * (b - y)
		mx = Min(mx, x)
		my = Min(my, y)
		ma = Max(ma, a)
		mb = Max(mb, b)
	}
	if (ma-mx)*(mb-my) != area {
		return false
	}
	for p, c := range freq {
		x, y := p[0], p[1]
		isCorner := (x == mx || x == ma) && (y == my || y == mb)
		if isCorner && c > 1 || !isCorner && c&1 == 1 {
			return false
		}
	}
	return true
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
