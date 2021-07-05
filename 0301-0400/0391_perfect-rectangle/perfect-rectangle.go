package main

// Tags:
// Math
// Hash

func isRectangleCover(rectangles [][]int) bool {
	sumA := 0
	minX, minY, maxA, maxB := 1<<31-1, 1<<31-1, -1<<31, -1<<31
	freq := map[[2]int]int{}

	for _, rec := range rectangles {
		x, y, a, b := rec[0], rec[1], rec[2], rec[3]
		freq[[2]int{x, y}]++
		freq[[2]int{x, b}]++
		freq[[2]int{a, y}]++
		freq[[2]int{a, b}]++
		minX = Min(minX, x)
		minY = Min(minY, y)
		maxA = Max(maxA, a)
		maxB = Max(maxB, b)
		sumA += (a - x) * (b - y)
	}

	sumB := (maxA - minX) * (maxB - minY)

	if sumA != sumB {
		return false
	}

	for p, c := range freq {
		x, y := p[0], p[1]
		isCorner := (x == minX || x == maxA) && (y == minY || y == maxB)
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
