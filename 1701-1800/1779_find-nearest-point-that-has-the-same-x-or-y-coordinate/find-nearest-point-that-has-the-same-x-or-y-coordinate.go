package main

// Difficulty:
// Easy

func nearestValidPoint(x int, y int, points [][]int) int {
	o, v := -1, 1<<31-1
	for i, e := range points {
		if e[0] != x && e[1] != y {
			continue
		}
		if d := abs(e[0]-x) + abs(e[1]-y); d < v {
			o, v = i, d
		}
	}
	return o
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
