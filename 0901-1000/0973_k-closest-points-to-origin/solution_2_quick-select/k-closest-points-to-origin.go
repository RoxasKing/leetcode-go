package main

import "math/rand"

// Difficulty:
// Medium

// Tags:
// Quick Select

func kClosest(points [][]int, K int) [][]int {
	quickSelect(points, 0, len(points)-1, K)
	return points[:K]
}

func quickSelect(points [][]int, l, r int, k int) {
	if l == r {
		return
	}
	pivotIndex := l + rand.Intn(r+1-l)
	points[pivotIndex], points[r] = points[r], points[pivotIndex]
	index := l
	for i := l; i < r; i++ {
		if isCloser(points[i], points[r]) {
			points[i], points[index] = points[index], points[i]
			index++
		}
	}
	points[index], points[r] = points[r], points[index]
	if index < k-1 {
		quickSelect(points, index+1, r, k)
	} else if index > k-1 {
		quickSelect(points, l, index-1, k)
	}
}

func isCloser(a, b []int) bool { return a[0]*a[0]+a[1]*a[1] < b[0]*b[0]+b[1]*b[1] }
