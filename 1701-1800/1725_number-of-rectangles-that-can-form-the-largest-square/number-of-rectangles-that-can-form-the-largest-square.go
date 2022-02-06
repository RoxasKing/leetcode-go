package main

// Difficulty:
// Easy

func countGoodRectangles(rectangles [][]int) int {
	max, out := 0, 0
	for _, r := range rectangles {
		if x := Min(r[0], r[1]); x > max {
			max, out = x, 1
		} else if x == max {
			out++
		}
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
