package main

import "sort"

// Difficulty:
// Hard

// Tags:
// Discretization
// Line Sweep
// Sorting
// Hash

func rectangleArea(rectangles [][]int) int {
	set := map[int]struct{}{}
	for _, x := range rectangles {
		set[x[1]] = struct{}{}
		set[x[3]] = struct{}{}
	}
	points := make([]int, 0, len(set))
	for p := range set {
		points = append(points, p)
	}
	sort.Ints(points)
	o := 0
	for k := 0; k < len(points)-1; k++ {
		lo, hi := points[k], points[k+1]
		a := []int{}
		for i, x := range rectangles {
			if x[1] <= lo && hi <= x[3] {
				a = append(a, i)
			}
		}
		if len(a) == 0 {
			continue
		}
		sort.Slice(a, func(i, j int) bool { return rectangles[a[i]][0] < rectangles[a[j]][0] })
		l, r := rectangles[a[0]][0], rectangles[a[0]][2]
		for _, i := range a[1:] {
			if r < rectangles[i][0] {
				o += (r - l) * (hi - lo)
				l = rectangles[i][0]
			}
			if r < rectangles[i][2] {
				r = rectangles[i][2]
			}
		}
		o += (r - l) * (hi - lo)
	}
	return o % (1e9 + 7)
}
