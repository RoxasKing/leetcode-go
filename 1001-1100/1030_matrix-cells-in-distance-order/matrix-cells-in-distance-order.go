package main

import "sort"

// Bucket Sort
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	maxDist := Max(r0, R-1-r0) + Max(c0, C-1-c0)
	buckets := make([][][]int, maxDist+1)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			dist := Abs(i-r0) + Abs(j-c0)
			buckets[dist] = append(buckets[dist], []int{i, j})
		}
	}
	out := make([][]int, 0, R*C)
	for _, b := range buckets {
		out = append(out, b...)
	}
	return out
}

// Sort
func allCellsDistOrder2(R int, C int, r0 int, c0 int) [][]int {
	out := make([][]int, 0, R*C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			out = append(out, []int{i, j})
		}
	}
	sort.Slice(out, func(i, j int) bool {
		return Abs(out[i][0]-r0)+Abs(out[i][1]-c0) < Abs(out[j][0]-r0)+Abs(out[j][1]-c0)
	})
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
