package main

import "math"

// Difficulty:
// Easy

// Tags:
// Math

func largestTriangleArea(points [][]int) float64 {
	cal := func(x1, y1, x2, y2, x3, y3 int) float64 {
		return math.Abs(float64(x1*y2+x2*y3+x3*y1-x1*y3-x2*y1-x3*y2)) / 2
	}
	var o float64
	for i, p := range points {
		for j, q := range points[:i] {
			for _, r := range points[:j] {
				o = math.Max(o, cal(p[0], p[1], q[0], q[1], r[0], r[1]))
			}
		}
	}
	return o
}
