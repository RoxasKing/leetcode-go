package main

import "sort"

// Discretization
// Line Sweep

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

	var out int
	for k := 0; k < len(points)-1; k++ {
		h1, h2 := points[k], points[k+1]
		h := h2 - h1
		idxs := []int{}
		for i, x := range rectangles {
			if x[1] <= h1 && h2 <= x[3] {
				idxs = append(idxs, i)
			}
		}
		if len(idxs) == 0 {
			continue
		}
		sort.Slice(idxs, func(i, j int) bool {
			i, j = idxs[i], idxs[j]
			return rectangles[i][0] < rectangles[j][0]
		})
		l, r := rectangles[idxs[0]][0], rectangles[idxs[0]][2]
		for _, i := range idxs[1:] {
			if rectangles[i][0] > r {
				out += (r - l) * h
				l, r = rectangles[i][0], rectangles[i][2]
			} else if rectangles[i][2] > r {
				r = rectangles[i][2]
			}
		}
		out += (r - l) * h
	}
	return out % (1e9 + 7)
}
