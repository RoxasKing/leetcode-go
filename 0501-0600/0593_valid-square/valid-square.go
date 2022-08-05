package main

import (
	"sort"
)

// Difficulty:
// Medium

// Tags:
// Math
// Sorting

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	arr := [][]int{p1, p2, p3, p4}
	sort.Slice(arr, func(i, j int) bool {
		a, b := arr[i], arr[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] < b[1]
	})
	cal := func(u, v []int) int { return (u[0]-v[0])*(u[0]-v[0]) + (u[1]-v[1])*(u[1]-v[1]) }
	a, b, c, d := arr[0], arr[1], arr[2], arr[3]
	ab := cal(a, b)
	cd := cal(c, d)
	ac := cal(a, c)
	bd := cal(b, d)
	bc := cal(b, c)
	da := cal(d, a)
	return ab > 0 && ab == cd && ac == bd && ab == ac && bc == da
}
