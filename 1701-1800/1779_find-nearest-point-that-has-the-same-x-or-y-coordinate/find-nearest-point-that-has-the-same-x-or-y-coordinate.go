package main

func nearestValidPoint(x int, y int, points [][]int) int {
	min, idx := 1<<31-1, -1
	for i, p := range points {
		if p[0] != x && p[1] != y {
			continue
		}
		dis := Abs(p[0]-x) + Abs(p[1]-y)
		if dis < min {
			min = dis
			idx = i
		}
	}
	return idx
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
