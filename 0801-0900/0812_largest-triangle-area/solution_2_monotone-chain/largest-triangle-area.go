package main

import (
	"math"
	"sort"
)

// Difficulty:
// Easy

// Tags:
// Monotone Chain

func largestTriangleArea(points [][]int) float64 {
	sort.Slice(points, func(i, j int) bool {
		a, b := points[i], points[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] < b[1]
	})
	arr := [][]int{}
	top := func() int { return len(arr) - 1 }
	cross := func(o, a, b []int) bool { return (o[1]-a[1])*(o[0]-b[0]) > (o[1]-b[1])*(o[0]-a[0]) }
	for _, p := range points {
		for ; len(arr) > 1 && cross(arr[top()-1], arr[top()], p); arr = arr[:top()] {
		}
		arr = append(arr, p)
	}
	for k, i := len(arr), len(points)-2; i >= 0; i-- {
		for ; len(arr) > k && cross(arr[top()-1], arr[top()], points[i]); arr = arr[:top()] {
		}
		arr = append(arr, points[i])
	}
	cal := func(x1, y1, x2, y2, x3, y3 int) float64 {
		return math.Abs(float64(x1*y2+x2*y3+x3*y1-x1*y3-x2*y1-x3*y2)) / 2
	}
	var o float64
	for i, a := range arr {
		for j, b := range arr[:i] {
			for _, c := range arr[:j] {
				o = math.Max(o, cal(a[0], a[1], b[0], b[1], c[0], c[1]))
			}
		}
	}
	return o
}
